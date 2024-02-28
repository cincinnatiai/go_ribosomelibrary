package stageservice

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
	"github.com/nicholaspark09/pipelineslibrary/models"
	"github.com/nicholaspark09/pipelineslibrary/models/request/stage"
	"log"
)

type StageServiceRepository struct {
	Endpoint       string
	ApiKey         string
	ContentType    string
	metricsManager metrics2.MetricsManagerContract
}

func ProvideStageRepository(endpoint string, apiKey string, metricsManger metrics2.MetricsManagerContract) StageServiceRepositoryContract {
	return &StageServiceRepository{
		Endpoint:       endpoint,
		ApiKey:         apiKey,
		ContentType:    "application/json",
		metricsManager: metricsManger,
	}
}

func (repo *StageServiceRepository) Create(request stage.StageCreateRequest) response.Response[models.Stage] {
	params := map[string]string{
		"action": "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Stage](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(request)
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 500
		return response.Response[models.Stage]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StageService.CreateStage"
	log.Printf("Trying to make a network call to Stages")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.metricsManager, func() (*models.Stage, *error) {
		callResponse, callError := network_v2.Post[models.Stage](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Stage]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[models.Stage]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[models.Stage]{Data: networkResponse, StatusCode: 200}
}

func (repo *StageServiceRepository) Fetch(request stage.FetchOneRequest) response.Response[models.Stage] {
	params := map[string]string{
		"partitionKey": request.PartitionKey,
		"rangeKey":     request.RangeKey,
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Stage](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	callName := "StageService.FetchOneStage"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.metricsManager, func() (*models.Stage, *error) {
		callResponse, err := network_v2.Get[models.Stage](manager)
		if err != nil {
			return nil, &err
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Stage]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[models.Stage]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[models.Stage]{Data: networkResponse, StatusCode: 200}
}

func (repo *StageServiceRepository) FetchAll(request stage.StageFetchAllRequest) response.Response[[]*models.Stage] {
	params := map[string]string{
		"action": "fetchAll",
	}
	manager := network_v2.ProvideNetworkManagerV2[*[]*models.Stage](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(request)
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[[]*models.Stage]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StageService.FetchAllStage"
	log.Printf("Trying to make a network call to Stages")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.metricsManager, func() (*[]*models.Stage, *error) {
		callResponse, callError := network_v2.Post[*[]*models.Stage](manager, bytes)
		if callError != nil {
			log.Printf("Got an error in calling stages: %s", callError.Error())
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		log.Printf("Error found in getting stages")
		return response.Response[[]*models.Stage]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		log.Printf("No data found")
		stages := make([]*models.Stage, 0)
		return response.Response[[]*models.Stage]{Data: &stages, StatusCode: statusCode, Message: "No stages found"}
	}
	return response.Response[[]*models.Stage]{Data: networkResponse, StatusCode: 200}
}

func (repo *StageServiceRepository) Update(request stage.UpdateRequest) response.Response[bool] {
	params := map[string]string{
		"action": "update",
	}
	manager := network_v2.ProvideNetworkManagerV2[*bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(request.Stage)
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StageService.UpdateStage"
	log.Printf("Trying to make a network call to Stages")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.metricsManager, func() (*bool, *error) {
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

func (repo *StageServiceRepository) Delete(request stage.DeleteRequest) response.Response[bool] {
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
		statusCode = 400
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StageService.DeleteStage"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.metricsManager, func() (*bool, *error) {
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
