package stageservice

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	"github.com/cincinnatiai/go_ribosomelibrary/models/request/stage"
	response "github.com/nicholaspark09/awsgorocket/model"
)

type StageServiceRepositoryContract interface {
	Create(request stage.StageCreateRequest) response.Response[models.Stage]
	Fetch(request stage.FetchOneRequest) response.Response[models.Stage]
	FetchAll(request stage.StageFetchAllRequest) response.Response[[]*models.Stage]
	Update(request stage.UpdateRequest) response.Response[bool]
	Delete(request stage.DeleteRequest) response.Response[bool]
}
