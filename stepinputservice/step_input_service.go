package stepinputservice

import (
	"encoding/json"
	"errors"
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	"github.com/cincinnatiai/go_ribosomelibrary/models/request"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
	"log"
)

type StepInputServiceContract interface {
	Create(campaignKey string, campaignRangeKey string, stepId string, input string, userId string, status string) response.Response[models.StepInput]
	Update(model models.StepInput) response.Response[bool]
	Delete(partitionKey string, rangeKey string, isHardDelete bool) response.Response[bool]
	Fetch(partitionKey string, rangeKey string) response.Response[models.StepInput]
	FetchAllByUser(identityId string, lastRangeKey *string) response.Response[[]*models.StepInput]
}

type StepInputService struct {
	Endpoint       string
	ApiKey         string
	ClientId       string
	ClientKey      string
	ContentType    string
	Controller     string
	MetricsManager metrics2.MetricsManagerContract
}

func (s StepInputService) Create(campaignKey string, campaignRangeKey string, stepId string, input string, userId string, status string) response.Response[models.StepInput] {
	createRequest := request.StepInputCreateRequest{
		ClientId:             s.ClientId,
		ClientKey:            s.ClientKey,
		CampaignPartitionKey: campaignKey,
		CampaignRangeKey:     campaignRangeKey,
		StepId:               stepId,
		Input:                input,
		UserId:               userId,
		Status:               status,
	}
	params := map[string]string{
		"controller": s.Controller,
		"action":     "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[models.StepInput](s.Endpoint, params, &s.ApiKey, &s.ContentType)
	bytes, err := json.Marshal(createRequest)
	if err != nil {
		log.Printf("StepInputService.CreateFailure: Failed to parse input: %s", err)
		return response.Response[models.StepInput]{Data: nil, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("StepInputService.Create", s.MetricsManager, func() (*models.StepInput, *error) {
		callResponse, callError := network_v2.Post[models.StepInput](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.StepInput]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[models.StepInput]{Data: nil, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[models.StepInput]{Data: networkResponse, StatusCode: 200}
}

func (s StepInputService) Update(model models.StepInput) response.Response[bool] {
	params := map[string]string{
		"controller": s.Controller,
		"action":     "update",
		"clientId":   s.ClientId,
	}
	manager := network_v2.ProvideNetworkManagerV2[bool](s.Endpoint, params, &s.ApiKey, &s.ContentType)
	bytes, err := json.Marshal(model)
	var result = false
	if err != nil {
		log.Printf("StepInputService.UpdateFailure: Failed to parse input: %s", err)
		return response.Response[bool]{Data: &result, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("StepInputService.Update", s.MetricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: &result, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[bool]{Data: &result, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (s StepInputService) Delete(partitionKey string, rangeKey string, isHardDelete bool) response.Response[bool] {
	params := map[string]string{
		"controller": s.Controller,
		"action":     "delete",
	}
	deleteRequest := request.StepInputDeleteRequest{
		ClientId:     s.ClientId,
		ClientKey:    s.ClientKey,
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		IsHardDelete: isHardDelete,
	}
	manager := network_v2.ProvideNetworkManagerV2[bool](s.Endpoint, params, &s.ApiKey, &s.ContentType)
	bytes, err := json.Marshal(deleteRequest)
	var result = false
	if err != nil {
		log.Printf("StepInputService.DeleteFailure: Failed to parse input: %s", err)
		return response.Response[bool]{Data: &result, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("StepInputService.Delete", s.MetricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: &result, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[bool]{Data: &result, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (s StepInputService) Fetch(partitionKey string, rangeKey string) response.Response[models.StepInput] {
	params := map[string]string{
		"controller":   s.Controller,
		"partitionKey": partitionKey,
		"rangeKey":     rangeKey,
	}
	manager := network_v2.ProvideNetworkManagerV2[models.StepInput](s.Endpoint, params, &s.ApiKey, &s.ContentType)
	callName := "StepInputService.FetchOne"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, s.MetricsManager, func() (*models.StepInput, *error) {
		callResponse, err := network_v2.Get[models.StepInput](manager)
		if err != nil {
			return nil, &err
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.StepInput]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[models.StepInput]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[models.StepInput]{Data: networkResponse, StatusCode: 200}
}

func (s StepInputService) FetchAllByUser(identityId string, lastRangeKey *string) response.Response[[]*models.StepInput] {
	if len(identityId) == 0 {
		return response.Response[[]*models.StepInput]{Data: nil, StatusCode: 401}
	}
	params := map[string]string{
		"controller": s.Controller,
		"action":     "fetchByUser",
	}
	fetchRequest := request.StepInputFetchByUserRequest{
		ClientId:  s.ClientId,
		ClientKey: s.ClientKey,
		UserId:    identityId,
	}
	bytes, err := json.Marshal(fetchRequest)
	if err != nil {
		log.Printf("StepInputService.DeleteFailure: Failed to parse input: %s", err)
		return response.Response[[]*models.StepInput]{Data: nil, StatusCode: 400, Message: "Failed to parse request"}
	}
	manager := network_v2.ProvideNetworkManagerV2[[]*models.StepInput](s.Endpoint, params, &s.ApiKey, &s.ContentType)
	callName := "StepInputService.FetchByUser"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, s.MetricsManager, func() (*[]*models.StepInput, *error) {
		callResponse, callErr := network_v2.Post[[]*models.StepInput](manager, bytes)
		if callErr != nil {
			return nil, &callErr
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[[]*models.StepInput]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[[]*models.StepInput]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[[]*models.StepInput]{Data: networkResponse, StatusCode: 200}
}

func ProvideStepInputService(
	endpoint string,
	apiKey string,
	clientId string,
	clientKey string,
	metricsManger metrics2.MetricsManagerContract) StepInputServiceContract {
	return &StepInputService{
		Endpoint:       endpoint,
		ApiKey:         apiKey,
		ClientId:       clientId,
		ClientKey:      clientKey,
		ContentType:    "application/json",
		Controller:     "stepinput",
		MetricsManager: metricsManger,
	}
}
