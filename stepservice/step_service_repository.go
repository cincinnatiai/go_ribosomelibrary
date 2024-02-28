package stepservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
	"github.com/nicholaspark09/pipelineslibrary/models"
	request "github.com/nicholaspark09/pipelineslibrary/models/request/step"
	response2 "github.com/nicholaspark09/pipelineslibrary/models/response/step"
	"log"
)

type StepServiceRepository struct {
	Endpoint       string
	ApiKey         string
	ContentType    string
	metricsManager metrics.MetricsManagerContract
}

func ProvideStepRepository(pfsEndpoint string, apiKey string, metricsManger metrics.MetricsManagerContract) StepServiceRepositoryContract {
	return &StepServiceRepository{
		Endpoint:       pfsEndpoint,
		ApiKey:         apiKey,
		ContentType:    "application/json",
		metricsManager: metricsManger,
	}
}

func (repo *StepServiceRepository) Create(request request.StepCreateRequest) response.Response[models.Step] {
	params := map[string]string{
		"action": "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[*models.Step](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json.Marshal(request)
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[models.Step]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StepService.CreateStep"
	log.Printf("Trying to make a network call to Steps")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*models.Step, *error) {
		callResponse, callError := network_v2.Post[*models.Step](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return *callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Step]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[models.Step]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[models.Step]{Data: networkResponse, StatusCode: 200}
}

func (repo *StepServiceRepository) Fetch(request request.StepFetchOneRequest) response.Response[models.Step] {
	params := map[string]string{
		"partitionKey": request.PartitionKey,
		"rangeKey":     request.RangeKey,
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Step](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	callName := "StepService.FetchOneStep"
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*models.Step, *error) {
		callResponse, callError := network_v2.Get[models.Step](manager)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Step]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[models.Step]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[models.Step]{Data: networkResponse, StatusCode: 200}
}

func (repo *StepServiceRepository) FetchAll(request request.StepFetchRequest) response.Response[[]*models.Step] {
	params := map[string]string{
		"action": "fetchAll",
	}
	manager := network_v2.ProvideNetworkManagerV2[response2.StepFetchResponse](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json.Marshal(request)
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[[]*models.Step]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StepService.FetchAllStep"
	log.Printf("Trying to make a network call to Steps")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*response2.StepFetchResponse, *error) {
		callResponse, callError := network_v2.Post[response2.StepFetchResponse](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[[]*models.Step]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[[]*models.Step]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[[]*models.Step]{Data: &networkResponse.Results, StatusCode: 200}
}

func (repo *StepServiceRepository) Update(request request.StepUpdateRequest) response.Response[bool] {
	params := map[string]string{
		"action": "update",
	}
	manager := network_v2.ProvideNetworkManagerV2[bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json.Marshal(request.Step)
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StepService.UpdateStep"
	log.Printf("Trying to make a network call to Steps")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: &result, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[bool]{Data: &result, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (repo *StepServiceRepository) Delete(request request.StepDeleteRequest) response.Response[bool] {
	params := map[string]string{
		"action": "delete",
	}
	manager := network_v2.ProvideNetworkManagerV2[bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json.Marshal(request)
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "StepService.DeleteStep"
	log.Printf("Trying to make a network call to Steps")
	networkResponse, networkError := metrics.MeasureTimeWithError(callName, repo.metricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: &result, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[bool]{Data: &result, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}
