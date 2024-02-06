package request

type StageStatusFetchRequest struct {
	ClientId   string `json:"client_id"`
	ClientKey  string `json:"client_key"`
	CampaignId string `json:"campaign_id"`
	StageId    string `json:"stage_id"`
}
