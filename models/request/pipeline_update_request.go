package request

import "main/model"

type PipelineUpdateRequest struct {
	ClientId  string         `json:"client_id"`
	ClientKey string         `json:"client_key"`
	Pipeline  model.Pipeline `json:"pipeline"`
}
