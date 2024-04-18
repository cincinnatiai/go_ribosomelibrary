package service

import (
	"encoding/json"
	"errors"
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	request2 "github.com/cincinnatiai/go_ribosomelibrary/models/request/category_request"
	response2 "github.com/cincinnatiai/go_ribosomelibrary/models/response"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/awsgorocket/network_v2"
	"github.com/nicholaspark09/awsgorocket/utils"
	"log"
)

type CategoryServiceContract interface {
	Create(title string, description string, status string) response.Response[models.Category]
	Fetch(partitionKey string, rangeKey string) response.Response[models.Category]
	FetchAll(lastRangeKey *string) response.Response[response2.CategoryFetchResponse]
	Update(category models.Category) response.Response[bool]
	Delete(partitionKey string, rangeKey string, isHardDelete bool) response.Response[bool]
}

type CategoryService struct {
	Endpoint       string
	ApiKey         string
	ContentType    string
	ClientId       string
	ClientKey      string
	Controller     string
	MetricsManager metrics2.MetricsManagerContract
}

func (categoryService *CategoryService) Create(title string, description string, status string) response.Response[models.Category] {
	createRequest := request2.PFSCategoryCreateRequest{
		ClientId:    categoryService.ClientId,
		ClientKey:   categoryService.ClientKey,
		Title:       title,
		Description: description,
		Status:      status,
	}
	params := map[string]string{
		"controller": categoryService.Controller,
		"action":     "create",
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Category](categoryService.Endpoint, params, &categoryService.ApiKey, &categoryService.ContentType)
	bytes, err := json.Marshal(createRequest)
	if err != nil {
		log.Printf("CategoryService.CreateFailure: Failed to parse input: %s", err)
		return response.Response[models.Category]{Data: nil, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("PFSCategoryService.Create", categoryService.MetricsManager, func() (*models.Category, *error) {
		callResponse, callError := network_v2.Post[models.Category](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Category]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[models.Category]{Data: nil, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[models.Category]{Data: networkResponse, StatusCode: 200}
}

func (categoryService *CategoryService) Fetch(partitionKey string, rangeKey string) response.Response[models.Category] {
	params := map[string]string{
		"controller":   categoryService.Controller,
		"partitionKey": partitionKey,
		"rangeKey":     rangeKey,
		"clientId":     categoryService.ClientId,
		"clientKey":    categoryService.ClientKey,
	}
	manager := network_v2.ProvideNetworkManagerV2[models.Category](categoryService.Endpoint, params, &categoryService.ApiKey, &categoryService.ContentType)
	callName := "CategoryService.FetchOne"
	networkResponse, networkError := metrics2.MeasureTimeWithError(callName, categoryService.MetricsManager, func() (*models.Category, *error) {
		callResponse, err := network_v2.Get[models.Category](manager)
		if err != nil {
			return nil, &err
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if networkError != nil && errors.As(*networkError, &genericError) {
		return response.Response[models.Category]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		errorMessage := "Error in making network call"
		return response.Response[models.Category]{Data: nil, StatusCode: 500, Message: errorMessage}
	}
	return response.Response[models.Category]{Data: networkResponse, StatusCode: 200}
}

func (categoryService *CategoryService) FetchAll(lastRangeKey *string) response.Response[response2.CategoryFetchResponse] {
	fetchRequest := request2.PFSCategoryFetchAllRequest{
		ClientId:     categoryService.ClientId,
		ClientKey:    categoryService.ClientKey,
		LastRangeKey: lastRangeKey,
	}
	params := map[string]string{
		"controller": categoryService.Controller,
		"action":     "fetchAll",
	}
	manager := network_v2.ProvideNetworkManagerV2[response2.CategoryFetchResponse](categoryService.Endpoint, params, &categoryService.ApiKey, &categoryService.ContentType)
	bytes, err := json.Marshal(fetchRequest)
	if err != nil {
		log.Printf("CategoryService.FetchAllFailure: Failed to parse input: %s", err)
		return response.Response[response2.CategoryFetchResponse]{Data: nil, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("PFSCategoryService.FetchAll", categoryService.MetricsManager, func() (*response2.CategoryFetchResponse, *error) {
		callResponse, callError := network_v2.Post[response2.CategoryFetchResponse](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[response2.CategoryFetchResponse]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[response2.CategoryFetchResponse]{Data: nil, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[response2.CategoryFetchResponse]{Data: networkResponse, StatusCode: 200}
}

func (categoryService *CategoryService) Update(category models.Category) response.Response[bool] {
	createRequest := request2.PFSCategoryUpdateRequest{
		ClientId:  categoryService.ClientId,
		ClientKey: categoryService.ClientKey,
		Category:  category,
	}
	params := map[string]string{
		"controller": categoryService.Controller,
		"action":     "update",
	}
	manager := network_v2.ProvideNetworkManagerV2[bool](categoryService.Endpoint, params, &categoryService.ApiKey, &categoryService.ContentType)
	bytes, err := json.Marshal(createRequest)
	if err != nil {
		log.Printf("CategoryService.UpdateFailure: Failed to parse input: %s", err)
		return response.Response[bool]{Data: nil, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("PFSCategoryService.Update", categoryService.MetricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[bool]{Data: nil, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func (categoryService *CategoryService) Delete(partitionKey string, rangeKey string, isHardDelete bool) response.Response[bool] {
	createRequest := request2.PFSCategoryDeleteRequest{
		ClientId:     categoryService.ClientId,
		ClientKey:    categoryService.ClientKey,
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		IsHardDelete: isHardDelete,
	}
	params := map[string]string{
		"controller": categoryService.Controller,
		"action":     "delete",
	}
	manager := network_v2.ProvideNetworkManagerV2[bool](categoryService.Endpoint, params, &categoryService.ApiKey, &categoryService.ContentType)
	bytes, err := json.Marshal(createRequest)
	if err != nil {
		log.Printf("CategoryService.DeleteFailure: Failed to parse input: %s", err)
		return response.Response[bool]{Data: nil, StatusCode: 400, Message: "Failed to parse request"}
	}
	networkResponse, networkError := metrics2.MeasureTimeWithError("PFSCategoryService.Delete", categoryService.MetricsManager, func() (*bool, *error) {
		callResponse, callError := network_v2.Post[bool](manager, bytes)
		if callError != nil {
			return nil, &callError
		}
		return callResponse, nil
	})
	var genericError utils.GenericError
	if err != nil && errors.As(*networkError, &genericError) {
		return response.Response[bool]{Data: nil, StatusCode: genericError.StatusCode, Message: genericError.Message}
	}
	if networkResponse == nil {
		return response.Response[bool]{Data: nil, StatusCode: 500, Message: "Error in making network call"}
	}
	return response.Response[bool]{Data: networkResponse, StatusCode: 200}
}

func ProvideCategoryService(
	endpoint string,
	apiKey string,
	clientId string,
	clientKey string,
	metricsManger metrics2.MetricsManagerContract) CategoryServiceContract {
	return &CategoryService{
		Endpoint:       endpoint,
		ApiKey:         apiKey,
		ClientId:       clientId,
		ClientKey:      clientKey,
		ContentType:    "application/json",
		Controller:     "category",
		MetricsManager: metricsManger,
	}
}
