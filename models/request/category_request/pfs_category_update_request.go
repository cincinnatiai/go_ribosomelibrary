package request

import "main/model"

type PFSCategoryUpdateRequest struct {
	ClientId  string         `json:"client_id"`
	ClientKey string         `json:"client_key"`
	Category  model.Category `json:"category"`
}
