package response

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
)

type PipelineFetchResponse struct {
	Results      []*models.Pipeline `json:"results"`
	LastRangeKey *string            `json:"last_range_key"`
}
