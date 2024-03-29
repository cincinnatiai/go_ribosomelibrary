package campaign

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
)

type FetchAllResponse struct {
	Results      []*models.Campaign `json:"results"`
	LastRangeKey *string            `json:"last_range_key"`
}
