package models

type StepInput struct {
	// StepId
	PartitionKey string `json:"partition_key"`
	// Random
	RangeKey             string `json:"range_key"`
	Input                string `json:"input"`
	Comments             string `json:"comments"`
	Created              string `json:"created"`
	Modified             string `json:"modified"`
	UserId               string `json:"user_id"`
	Status               string `json:"status"`
	CampaignPartitionKey string `json:"campaign_partition_key"`
	ApprovalType         string `json:"approval_type"`
}
