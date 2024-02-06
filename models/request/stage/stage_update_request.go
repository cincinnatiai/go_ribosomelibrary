package stage

import "main/model"

type UpdateRequest struct {
	ClientId  string      `json:"client_id"`
	ClientKey string      `json:"client_key"`
	Stage     model.Stage `json:"stage"`
}
