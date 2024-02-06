package campaign

import "main/model"

type UpdateRequest struct {
	ClientId  string         `json:"client_id"`
	ClientKey string         `json:"client_key"`
	Campaign  model.Campaign `json:"campaign"`
}
