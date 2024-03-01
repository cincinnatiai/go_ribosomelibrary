package service

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	"github.com/cincinnatiai/go_ribosomelibrary/models/request/campaign"
	campaign2 "github.com/cincinnatiai/go_ribosomelibrary/models/response/campaign"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
	"log"
)

type CampaignService struct {
	Endpoint       string
	ApiKey         string
	ClientId       string
	ClientKey      string
	ContentType    string
	Controller     string
	MetricsManager metrics2.MetricsManagerContract
}

func ProvideCampaignService(
	endpoint string,
	apiKey string,
	clientId string,
	clientKey string,
	metricsManger metrics2.MetricsManagerContract) CampaignServiceContract {
	return &CampaignService{
		Endpoint:       endpoint,
		ApiKey:         apiKey,
		ClientId:       clientId,
		ClientKey:      clientKey,
		ContentType:    "application/json",
		Controller:     "campaigns",
		MetricsManager: metricsManger,
	}
}

func (repo *CampaignService) Create(pipelinePartitionKey string, pipelineRangeKey string, title string, description string, secondaryId string) response.Response[models.Campaign] {
	params := map[string]string{
		"controller": repo.Controller,
		"action":     "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Campaign](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(campaign.CreateRequest{
		ClientId:             repo.ClientId,
		ClientKey:            repo.ClientKey,
		PipelinePartitionKey: pipelinePartitionKey,
		PipelineRangeKey:     pipelineRangeKey,
		Title:                title,
		Description:          description,
		SecondaryId:          secondaryId,
	})
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 500
		return response.Response[models.Campaign]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "CampaignService.Create"
	log.Printf("Trying to make a network call to campaign")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.MetricsManager, func() (*models.Campaign, *error) {
		callResponse, callError := network_v2.Post[models.Campaign](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Campaign]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		statusCode = 500
		errorMessage := "Error in making network call"
		return response.Response[models.Campaign]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	return response.Response[models.Campaign]{Data: networkResponse, StatusCode: 200}
}

func (repo *CampaignService) FetchAllByUser(identityId string) response.Response[[]*models.Campaign] {
	params := map[string]string{
		"controller": "campaigns",
		"action":     "fetchByUser",
	}
	manager := network_v2.ProvideNetworkManagerV2[[]*models.Campaign](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(campaign.FetchByUser{
		ClientId:  repo.ClientId,
		ClientKey: repo.ClientKey,
		UserId:    identityId,
	})
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[[]*models.Campaign]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "CampaignService.FetchAllByUser"
	log.Printf("Trying to make a network call to campaigns")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.MetricsManager, func() (*[]*models.Campaign, *error) {
		callResponse, callError := network_v2.Post[[]*models.Campaign](manager, bytes)
		if callError != nil {
			log.Printf("Got an error in calling campaigns: %s", callError.Error())
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		log.Printf("Error found in getting campaigns")
		return response.Response[[]*models.Campaign]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		log.Printf("No data found")
		campaigns := make([]*models.Campaign, 0)
		return response.Response[[]*models.Campaign]{Data: &campaigns, StatusCode: statusCode, Message: "No campaigns found"}
	}
	return response.Response[[]*models.Campaign]{Data: networkResponse, StatusCode: 200}
}

func (repo *CampaignService) Fetch(partitionKey string, rangeKey string) response.Response[models.Campaign] {
	params := map[string]string{
		"controller":   "campaigns",
		"partitionKey": partitionKey,
		"rangeKey":     rangeKey,
		"clientId":     repo.ClientId,
		"clientKey":    repo.ClientKey,
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Campaign](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	callName := "CampaignService.FetchOne"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.MetricsManager, func() (*models.Campaign, *error) {
		callResponse, err := network_v2.Get[models.Campaign](manager)
		if err != nil {
			return nil, &err
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Campaign]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[models.Campaign]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[models.Campaign]{Data: networkResponse, StatusCode: 200}
}

func (repo *CampaignService) FetchAll(pipelinePartitionKey string, pipelineRangeKey string, lastRangeKey *string) response.Response[campaign2.FetchAllResponse] {
	params := map[string]string{
		"controller": "campaigns",
		"action":     "fetchAll",
	}
	manager := network_v2.ProvideNetworkManagerV2[campaign2.FetchAllResponse](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(campaign.FetchAllRequest{
		ClientId:             repo.ClientId,
		ClientKey:            repo.ClientKey,
		PipelinePartitionKey: pipelinePartitionKey,
		PipelineRangeKey:     pipelineRangeKey,
		LastRangeKey:         lastRangeKey,
	})
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[campaign2.FetchAllResponse]{Data: nil, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "CampaignService.FetchAll"
	log.Printf("Trying to make a network call to s")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.MetricsManager, func() (*campaign2.FetchAllResponse, *error) {
		callResponse, callError := network_v2.Post[campaign2.FetchAllResponse](manager, bytes)
		if callError != nil {
			log.Printf("Got an error in calling campaigns: %s", callError.Error())
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		log.Printf("Error found in getting campaigns")
		return response.Response[campaign2.FetchAllResponse]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		log.Printf("No data found")
		campaigns := make([]*models.Campaign, 0)
		return response.Response[campaign2.FetchAllResponse]{Data: &campaign2.FetchAllResponse{Results: campaigns}, StatusCode: statusCode, Message: "No campaigns found"}
	}
	return response.Response[campaign2.FetchAllResponse]{Data: networkResponse, StatusCode: 200}
}

func (repo *CampaignService) Update(model models.Campaign) response.Response[bool] {
	params := map[string]string{
		"controller": "campaigns",
		"action":     "update",
	}
	manager := network_v2.ProvideNetworkManagerV2[*bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(campaign.UpdateRequest{
		ClientId:  repo.ClientId,
		ClientKey: repo.ClientKey,
		Campaign:  model,
	})
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "CampaignService.Update"
	log.Printf("Trying to make a network call to Campaigns")
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.MetricsManager, func() (*bool, *error) {
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

func (repo *CampaignService) Delete(partitionKey string, rangeKey string, isHardDelete bool) response.Response[bool] {
	params := map[string]string{
		"controller": "campaigns",
		"action":     "delete",
	}
	manager := network_v2.ProvideNetworkManagerV2[*bool](repo.Endpoint, params, &repo.ApiKey, &repo.ContentType)
	bytes, err := json2.Marshal(campaign.DeleteRequest{
		ClientId:     repo.ClientId,
		ClientKey:    repo.ClientKey,
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		IsHardDelete: isHardDelete,
	})
	var result = false
	var statusCode int
	if err != nil {
		log.Printf("Error in converting request to json: %s", err.Error())
		errorMessage := fmt.Sprintf("Error in converting request to json: %s", err.Error())
		statusCode = 400
		return response.Response[bool]{Data: &result, StatusCode: statusCode, Message: errorMessage}
	}
	callName := "CampaignService.Delete"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, repo.MetricsManager, func() (*bool, *error) {
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
