package stepservice

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	request "github.com/cincinnatiai/go_ribosomelibrary/models/request/step"
	response "github.com/nicholaspark09/awsgorocket/model"
)

type StepServiceRepositoryContract interface {
	Create(request request.StepCreateRequest) response.Response[models.Step]
	Fetch(request request.StepFetchOneRequest) response.Response[models.Step]
	FetchAll(request request.StepFetchRequest) response.Response[[]*models.Step]
	Update(request request.StepUpdateRequest) response.Response[bool]
	Delete(request request.StepDeleteRequest) response.Response[bool]
}
