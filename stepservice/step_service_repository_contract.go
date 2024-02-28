package stepservice

import (
	response "github.com/nicholaspark09/awsgorocket/model"
	"github.com/nicholaspark09/pipelineslibrary/models"
	request "github.com/nicholaspark09/pipelineslibrary/models/request/step"
)

type StepServiceRepositoryContract interface {
	Create(request request.StepCreateRequest) response.Response[models.Step]
	Fetch(request request.StepFetchOneRequest) response.Response[models.Step]
	FetchAll(request request.StepFetchRequest) response.Response[[]*models.Step]
	Update(request request.StepUpdateRequest) response.Response[bool]
	Delete(request request.StepDeleteRequest) response.Response[bool]
}
