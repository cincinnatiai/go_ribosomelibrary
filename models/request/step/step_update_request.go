package request

import "github.com/nicholaspark09/pipelineslibrary/models"

type StepUpdateRequest struct {
	ClientId  string      `json:"client_id"`
	ClientKey string      `json:"client_key"`
	Step      models.Step `json:"step"`
}
