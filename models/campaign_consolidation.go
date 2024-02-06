package models

type CampaignConsolidation struct {
	Campaign *Campaign              `json:"campaign"`
	Pipeline *PipelineWithAllModels `json:"pipeline"`
}
