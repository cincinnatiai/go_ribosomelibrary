package campaign

import "github.com/nicholaspark09/pipelineslibrary/models"

type UpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	Campaign  models.Campaign `json:"campaign"`
}
