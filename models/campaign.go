package models

type Campaign struct {
	// Should be the clientId_pipelineId
	PartitionKey string `json:"partition_key"`
	// Time.RandomUUID
	RangeKey        string `json:"range_key"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	CurrentStageId  string `json:"current_stage_id"`
	CurrentStatusId string `json:"current_status_id"`
	Created         string `json:"created"`
	Modified        string `json:"modified"`
	Status          string `json:"status"`
	UserId          string `json:"user_id"`
	DueBy           string `json:"due_by"`
}
