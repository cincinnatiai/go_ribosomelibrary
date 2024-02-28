package request

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
)

type PipelineUpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	Pipeline  models.Pipeline `json:"pipeline"`
}
