package stageservice

import (
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/pipelineslibrary/models"
	"github.com/nicholaspark09/pipelineslibrary/models/request/stage"
)

type StageServiceRepositoryContract interface {
	Create(request stage.StageCreateRequest) response.Response[models.Stage]
	Fetch(request stage.FetchOneRequest) response.Response[models.Stage]
	FetchAll(request stage.StageFetchAllRequest) response.Response[[]*models.Stage]
	Update(request stage.UpdateRequest) response.Response[bool]
	Delete(request stage.DeleteRequest) response.Response[bool]
}
