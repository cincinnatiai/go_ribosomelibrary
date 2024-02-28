package stage

import "github.com/nicholaspark09/pipelineslibrary/models"

type StageFetchResponse struct {
	Results []*models.Stage `json:"results"`
}
