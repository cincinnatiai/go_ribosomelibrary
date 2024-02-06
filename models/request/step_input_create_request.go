package request

type StepInputCreateRequest struct {
	ClientId             string `json:"client_id"`
	ClientKey            string `json:"client_key"`
	CampaignPartitionKey string `json:"campaign_partition_key"`
	CampaignRangeKey     string `json:"campaign_range_key"`
	StepId               string `json:"step_id"`
	Input                string `json:"input"`
	UserId               string `json:"user_id"`
	Status               string `json:"status"`
}
