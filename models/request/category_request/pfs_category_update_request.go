package request

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type PFSCategoryUpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	Category  models.Category `json:"category"`
}
