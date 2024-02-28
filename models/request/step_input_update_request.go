package request

import (
	"github.com/nicholaspark09/pipelineslibrary/models"
)

type StepInputUpdateRequest struct {
	ClientId  string           `json:"client_id"`
	ClientKey string           `json:"client_key"`
	StepInput models.StepInput `json:"step_input"`
}
