package request

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type StepUpdateRequest struct {
	ClientId  string      `json:"client_id"`
	ClientKey string      `json:"client_key"`
	Step      models.Step `json:"step"`
}
