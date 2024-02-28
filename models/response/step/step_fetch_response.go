package response

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type StepFetchResponse struct {
	Results      []*models.Step `json:"results"`
	LastRangeKey *string        `json:"last_range_key"`
}
