package request

import "github.com/nicholaspark09/pipelineslibrary/models"

type PFSCategoryUpdateRequest struct {
	ClientId  string          `json:"client_id"`
	ClientKey string          `json:"client_key"`
	Category  models.Category `json:"category"`
}
