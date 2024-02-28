package response

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type CategoryFetchResponse struct {
	Results      []*models.Category `json:"results"`
	LastRangeKey *string            `json:"last_range_key"`
}
