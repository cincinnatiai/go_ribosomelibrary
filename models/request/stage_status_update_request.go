package request

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
)

type StageStatusUpdateRequest struct {
	ClientId    string             `json:"client_id"`
	ClientKey   string             `json:"client_key"`
	StageStatus models.StageStatus `json:"stage_status"`
}
