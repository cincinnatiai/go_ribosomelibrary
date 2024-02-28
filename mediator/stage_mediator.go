package mediator

import (
	"fmt"
	"github.com/nicholaspark09/awsgorocket/utils"
	"github.com/nicholaspark09/pipelineslibrary/models"
	request2 "github.com/nicholaspark09/pipelineslibrary/models/request/step"
	stepservice2 "github.com/nicholaspark09/pipelineslibrary/stepservice"
	"log"
)

type StageMediatorContract interface {
	Fetch(
		clientId string,
		clientKey string,
		pipelineId string,
		stage *models.Stage) (*models.StageWithStep, error)
}

type StageMediator struct {
	service stepservice2.StepServiceRepositoryContract
}

func ProvideStageMediator(service stepservice2.StepServiceRepositoryContract) StageMediatorContract {
	return StageMediator{
		service: service,
	}
}

func (mediator StageMediator) Fetch(
	clientId string,
	clientKey string,
	pipelineId string,
	stage *models.Stage) (*models.StageWithStep, error) {
	log.Printf("Trying to fetch steps: %s", stage.Id)
	response := mediator.service.FetchAll(request2.StepFetchRequest{
		ClientId:   clientId,
		ClientKey:  clientKey,
		PipelineId: pipelineId,
		StageId:    stage.Id,
	})

	if response.StatusCode != 200 {
		return nil, utils.GenericError{Message: fmt.Sprintf("Error in fetching steps: %s", response.Message)}
	}
	var steps []*models.Step
	if *response.Data == nil || len(*response.Data) == 0 {
		steps = make([]*models.Step, 0)
	} else {
		steps = *response.Data
	}
	return &models.StageWithStep{
		Id:           stage.Id,
		PartitionKey: stage.PartitionKey,
		RangeKey:     stage.RangeKey,
		Title:        stage.Title,
		Description:  stage.Description,
		Body:         stage.Body,
		Modified:     stage.Modified,
		IsRequired:   stage.IsRequired,
		Type:         stage.Type,
		Status:       stage.Status,
		Steps:        steps,
	}, nil
}
