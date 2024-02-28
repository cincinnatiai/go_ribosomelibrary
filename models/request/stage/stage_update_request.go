package stage

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type UpdateRequest struct {
	ClientId  string       `json:"client_id"`
	ClientKey string       `json:"client_key"`
	Stage     models.Stage `json:"stage"`
}
