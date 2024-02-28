package campaign

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type UpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	Campaign  models.Campaign `json:"campaign"`
}
