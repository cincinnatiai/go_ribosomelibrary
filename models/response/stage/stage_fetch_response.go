package stage

import "main/model"

type StageFetchResponse struct {
	Results []*model.Stage `json:"results"`
}
