package pipelineservice

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/cincinnatiai/go_ribosomelibrary/mediator"
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	"github.com/cincinnatiai/go_ribosomelibrary/models/request"
	response2 "github.com/cincinnatiai/go_ribosomelibrary/models/response"
	"github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
	"log"
	"sort"
	"sync"
)

type PipelineServiceContract interface {
	CreatePipeline(request request.PipelineCreateRequest) response.Response[models.Pipeline]
	FetchSimplePipeline(oneRequest request.PipelineFetchOneRequest) response.Response[models.Pipeline]
	FetchPipeline(oneRequest request.PipelineFetchOneRequest) response.Response[models.PipelineWithAllModels]
	FetchAll(request request.PipelineFetchRequest) response.Response[response2.PipelineFetchResponse]
	UpdatePipeline(oneRequest request.PipelineUpdateRequest) response.Response[bool]
	DeletePipeline(request request.PipelineDeleteRequest) response.Response[bool]
}

type PipelineService struct {
	Endpoint       string
	ApiKey         string
	ContentType    string
	mediator       mediator.PipelineMediatorContract
	stageMediator  mediator.StageMediatorContract
	metricsManager metrics.MetricsManagerContract
	cache          map[string]*models.Pipeline
}

func ProvidePipelineService(
	endpoint string,
	apiKey string,
	pipelineMediator mediator.PipelineMediatorContract,
	stageMediator mediator.StageMediatorContract,
	metricsManger metrics.MetricsManagerContract) PipelineServiceContract {
	return &PipelineService{
		Endpoint:       endpoint,
		ApiKey:         apiKey,
		ContentType:    "application/json",
		mediator:       pipelineMediator,
		stageMediator:  stageMediator,
		metricsManager: metricsManger,
	}
}

func (repo *PipelineService) CreatePipeline(request request.PipelineCreateRequest) response.Response[models.Pipeline] {
	params := map[string]string{
		"action": "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[*models.Pipeline](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(request)
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 500
		return response.Response[models.Pipeline]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "PipelineService.CreatePipeline"
	log.Printf("Trying to make a network call to Pipelines")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*models.Pipeline, *error) {
		callResponse, callError := network_v2.Post[*models.Pipeline](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Pipeline]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[models.Pipeline]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[models.Pipeline]{Data: networkResponse, StatusCode: 200}
}

func (repo *PipelineService) FetchSimplePipeline(oneRequest request.PipelineFetchOneRequest) response.Response[models.Pipeline] {
	params := map[string]string{
		"partitionKey": oneRequest.PartitionKey,
		"rangeKey":     oneRequest.RangeKey,
	}
	cacheKey := buildCacheKey(oneRequest.PartitionKey, oneRequest.RangeKey)
	value, ok := repo.cache[cacheKey]
	if ok {
		return response.Response[models.Pipeline]{Data: value, StatusCode: 200}
	}
	manager := network_v2.ProvideNetworkManagerV2[*models.Pipeline](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	callName := "PipelineService.FetchOnePipeline"
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*models.Pipeline, *error) {
		callResponse, callError := network_v2.Get[*models.Pipeline](manager)
		if callError != nil {
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Pipeline]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[models.Pipeline]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[models.Pipeline]{Data: networkResponse, StatusCode: 200}
}

func (repo *PipelineService) FetchPipeline(oneRequest request.PipelineFetchOneRequest) response.Response[models.PipelineWithAllModels] {
	networkResponse := repo.FetchSimplePipeline(oneRequest)
	// Mediate the response
	consolidatedResponse, err := repo.mediator.Fetch(oneRequest.ClientId, oneRequest.ClientKey, *networkResponse.Data)
	var statusCode int
	if err != nil {
		statusCode = 500
		errorMessage := "Error in fetching and wrapping stages to Pipeline"
		log.Printf(errorMessage)
		return response.Response[models.PipelineWithAllModels]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	var stagesWithSteps []*models.StageWithStep
	if len(consolidatedResponse.Stages) > 0 {
		var waitingItems sync.WaitGroup
		for i, _ := range consolidatedResponse.Stages {
			waitingItems.Add(1)
			go func(s *models.Stage) {
				defer waitingItems.Done()
				log.Printf("Trying to make fetch Stage: %s, PipelineId: %s, stage", s.Id, consolidatedResponse.Id)
				mediatorResponse, stageError := repo.stageMediator.Fetch(oneRequest.ClientId, oneRequest.ClientKey, consolidatedResponse.Id, s)
				if stageError != nil {
					log.Printf("Error in fetching through stage mediator: %s", stageError.Error())
				} else if mediatorResponse != nil && &mediatorResponse != nil {
					sort.Slice(mediatorResponse.Steps, func(i, j int) bool {
						return mediatorResponse.Steps[i].RangeKey < mediatorResponse.Steps[j].RangeKey
					})
					stagesWithSteps = append(stagesWithSteps, mediatorResponse)
				}
			}(consolidatedResponse.Stages[i])
		}
		waitingItems.Wait()
	}
	if stagesWithSteps == nil {
		log.Printf("Stage with steps was nil")
		stagesWithSteps = make([]*models.StageWithStep, 0)
	} else {
		log.Printf("Stage with steps was not nil")
	}
	sort.Slice(stagesWithSteps, func(i, j int) bool {
		return stagesWithSteps[i].RangeKey < stagesWithSteps[j].RangeKey
	})
	pipelineWithAllModels := models.PipelineWithAllModels{
		Id:              consolidatedResponse.Id,
		PartitionKey:    consolidatedResponse.PartitionKey,
		RangeKey:        consolidatedResponse.RangeKey,
		Title:           consolidatedResponse.Title,
		Description:     consolidatedResponse.Description,
		Created:         consolidatedResponse.Created,
		Modified:        consolidatedResponse.Modified,
		Status:          consolidatedResponse.Status,
		IsPublic:        consolidatedResponse.IsPublic,
		AuxiliarHashKey: consolidatedResponse.AuxiliarHashKey,
		Stages:          stagesWithSteps,
		Type:            consolidatedResponse.Type,
	}
	return response.Response[models.PipelineWithAllModels]{StatusCode: 200, Data: &pipelineWithAllModels}
}

func (repo *PipelineService) FetchAll(request request.PipelineFetchRequest) response.Response[response2.PipelineFetchResponse] {
	params := map[string]string{
		"action": "fetchAll",
	}
	manager := network_v2.ProvideNetworkManagerV2[*response2.PipelineFetchResponse](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(request)
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 500
		return response.Response[response2.PipelineFetchResponse]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "PipelineService.FetchAllPipeline"
	log.Printf("Trying to make a network call to Pipelines")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*response2.PipelineFetchResponse, *error) {
		callResponse, callError := network_v2.Post[*response2.PipelineFetchResponse](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[response2.PipelineFetchResponse]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[response2.PipelineFetchResponse]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[response2.PipelineFetchResponse]{Data: networkResponse, StatusCode: 200}
}

func (repo *PipelineService) UpdatePipeline(oneRequest request.PipelineUpdateRequest) response.Response[bool] {
	params := map[string]string{
		"action": "update",
	}
	manager := network_v2.ProvideNetworkManagerV2[*bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(oneRequest)
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 500
		return response.Response[bool]{
			Data:       &result,
			StatusCode: statusCode,
			Message:    errorMessage,
		}
	}
	callName := "PipelineService.UpdatePipeline"
	log.Printf("Trying to make a network call to Pipelines")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[*bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: &result, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (repo *PipelineService) DeletePipeline(request request.PipelineDeleteRequest) response.Response[bool] {
	params := map[string]string{
		"action": "delete",
	}
	manager := network_v2.ProvideNetworkManagerV2[*bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(request)
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 500
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "PipelineService.DeletePipeline"
	log.Printf("Trying to make a network call to Pipelines")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[*bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: &result, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (repo *PipelineService) ProvidePipeline(clientId string, clientKey string, partitionKey string, rangeKey string) response.Response[models.Pipeline] {

	networkResponse := repo.FetchSimplePipeline(request.PipelineFetchOneRequest{
		ClientId:     clientId,
		ClientKey:    clientKey,
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
	})
	if networkResponse.StatusCode != 200 {
		log.Printf("Error in fetching pipeline: %v", networkResponse.StatusCode)
	}
	return networkResponse
}

func buildCacheKey(partitionKey string, rangeKey string) string {
	return partitionKey + "_" + rangeKey
}
