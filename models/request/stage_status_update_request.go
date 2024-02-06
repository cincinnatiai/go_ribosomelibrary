package request

import "main/model"

type StageStatusUpdateRequest struct {
	ClientId    string            `json:"client_id"`
	ClientKey   string            `json:"client_key"`
	StageStatus model.StageStatus `json:"stage_status"`
}
