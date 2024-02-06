package models

type StageStatus struct {
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
	Status       string `json:"status"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	ApprovalType string `json:"approval_type"`
	StepRangeKey string `json:"step_range_key"`
	UserId       string `json:"user_id"`
}
