package request

import (
	"github.com/nicholaspark09/pipelineslibrary/models"
)

type PipelineUpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	Pipeline  models.Pipeline `json:"pipeline"`
}
