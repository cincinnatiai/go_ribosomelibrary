package mediator

import (
	"github.com/nicholaspark09/awsgorocket/utils"
	"github.com/nicholaspark09/pipelineslibrary/models"
	"github.com/nicholaspark09/pipelineslibrary/models/request/stage"
	"github.com/nicholaspark09/pipelineslibrary/stageservice"
	"log"
	"sort"
)

type PipelineMediatorContract interface {
	Fetch(clientId string, clientKey string, pipeline models.Pipeline) (*models.PipelineConsolidation, error)
}

type PipelineMediator struct {
	service stageservice.StageServiceRepositoryContract
}

func ProvidePipelineMediator(service stageservice.StageServiceRepositoryContract) PipelineMediatorContract {
	return &PipelineMediator{
		service: service,
	}
}

func (mediator *PipelineMediator) Fetch(clientId string, clientKey string, pipeline models.Pipeline) (*models.PipelineConsolidation, error) {
	response := mediator.service.FetchAll(stage.StageFetchAllRequest{
		ClientId:   clientId,
		ClientKey:  clientKey,
		PipelineId: pipeline.Id,
	})
	if response.StatusCode == 200 {
		log.Printf("Got a 200 in fetchin stages")
		var stages []*models.Stage
		if *response.Data == nil {
			log.Printf("Data was nil")
			stages = make([]*models.Stage, 0)
		} else {
			log.Printf("Data was not nil")
			stages = *response.Data
		}
		sort.Slice(stages, func(i, j int) bool {
			return stages[i].RangeKey < stages[j].RangeKey
		})
		return &models.PipelineConsolidation{
			Id:              pipeline.Id,
			PartitionKey:    pipeline.PartitionKey,
			RangeKey:        pipeline.RangeKey,
			Title:           pipeline.Title,
			Description:     pipeline.Description,
			Created:         pipeline.Created,
			Modified:        pipeline.Modified,
			Status:          pipeline.Status,
			IsPublic:        pipeline.IsPublic,
			AuxiliarHashKey: pipeline.AuxiliarHashKey,
			Stages:          stages,
			Type:            pipeline.Type,
		}, nil
	} else {
		log.Printf("Got an error in trying to fetch stage")
		return nil, utils.GenericError{Message: response.Message}
	}
}
