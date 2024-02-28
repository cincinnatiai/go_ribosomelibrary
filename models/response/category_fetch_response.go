package response

import "github.com/nicholaspark09/pipelineslibrary/models"

type CategoryFetchResponse struct {
	Results      []*models.Category `json:"results"`
	LastRangeKey *string            `json:"last_range_key"`
}
