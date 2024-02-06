package request

import "main/model"

type StepInputUpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	StepInput model.StepInput `json:"step_input"`
}
