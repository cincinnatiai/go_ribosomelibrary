package request

type StageStatusCreateRequest struct {
	ClientId     string `json:"client_id"`
	ClientKey    string `json:"client_key"`
	CampaignId   string `json:"campaign_id"`
	StageId      string `json:"stage_id"`
	StepId       string `json:"step_id"`
	UserId       string `json:"user_id"`
	Status       string `json:"status"`
	ApprovalType string `json:"approval_type"`
}
