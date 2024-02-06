package request

import "main/model"

type StepUpdateRequest struct {
	ClientId  string     `json:"client_id"`
	ClientKey string     `json:"client_key"`
	Step      model.Step `json:"step"`
}
