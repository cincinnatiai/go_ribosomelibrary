package models

type StageConsolidation struct {
	Stage *Stage  `json:"stage"`
	Steps []*Step `json:"steps"`
}
