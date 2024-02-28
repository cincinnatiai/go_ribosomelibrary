package response

import (
	"github.com/nicholaspark09/pipelineslibrary/models"
)

type PipelineFetchResponse struct {
	Results      []*models.Pipeline `json:"results"`
	LastRangeKey *string            `json:"last_range_key"`
}
