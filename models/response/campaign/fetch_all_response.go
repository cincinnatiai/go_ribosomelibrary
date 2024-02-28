package campaign

import (
	"github.com/nicholaspark09/pipelineslibrary/models"
)

type FetchAllResponse struct {
	Results      []*models.Campaign `json:"results"`
	LastRangeKey *string            `json:"last_range_key"`
}
