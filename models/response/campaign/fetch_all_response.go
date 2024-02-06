package campaign

import "main/model"

type FetchAllResponse struct {
	Results      []*model.Campaign `json:"results"`
	LastRangeKey *string           `json:"last_range_key"`
}
