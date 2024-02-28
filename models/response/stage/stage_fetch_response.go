package stage

import "github.com/cincinnatiai/go_ribosomelibrary/models"

type StageFetchResponse struct {
	Results []*models.Stage `json:"results"`
}
