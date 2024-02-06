package models

type PipelineConsolidation struct {
	Id              string   `json:"id"`
	PartitionKey    string   `json:"partition_key"`
	RangeKey        string   `json:"range_key"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Created         string   `json:"created"`
	Modified        string   `json:"modified"`
	Status          string   `json:"status"`
	IsPublic        bool     `json:"is_public"`
	AuxiliarHashKey string   `json:"auxiliar_hash_key"`
	Type            *string  `json:"type"`
	Stages          []*Stage `json:"stages"`
}
