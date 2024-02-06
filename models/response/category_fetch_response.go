package response

import "main/model"

type CategoryFetchResponse struct {
	Results      []*model.Category `json:"results"`
	LastRangeKey *string           `json:"last_range_key"`
}
