package models

type CampaignWithStatus struct {
	// Should be the clientId_pipelineId
	PartitionKey string `json:"partition_key"`
	// Time.RandomUUID
	RangeKey       string         `json:"range_key"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	CurrentStageId string         `json:"current_stage_id"`
	CurrentStepId  string         `json:"current_step_id"`
	Created        string         `json:"created"`
	Modified       string         `json:"modified"`
	Status         string         `json:"status"`
	UserId         string         `json:"user_id"`
	DueBy          string         `json:"due_by"`
	StageStatuses  []*StageStatus `json:"stage_statuses"`
}

func Convert(campaign Campaign, statuses []*StageStatus) CampaignWithStatus {
	return CampaignWithStatus{
		PartitionKey:   campaign.PartitionKey,
		RangeKey:       campaign.RangeKey,
		Title:          campaign.Title,
		Description:    campaign.Description,
		CurrentStageId: campaign.CurrentStageId,
		Created:        campaign.Created,
		Modified:       campaign.Modified,
		Status:         campaign.Status,
		UserId:         campaign.UserId,
		DueBy:          campaign.DueBy,
		StageStatuses:  statuses,
	}
}
