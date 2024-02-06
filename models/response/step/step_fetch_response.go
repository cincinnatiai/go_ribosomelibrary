package response

import "main/model"

type StepFetchResponse struct {
	Results      []*model.Step `json:"results"`
	LastRangeKey *string       `json:"last_range_key"`
}
