package response

import "main/model"

type PipelineFetchResponse struct {
	Results      []*model.Pipeline `json:"results"`
	LastRangeKey *string           `json:"last_range_key"`
}
