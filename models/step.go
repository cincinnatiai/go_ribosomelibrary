package models

type Step struct {
	Id           string  `json:"id"`
	PartitionKey string  `json:"partition_key"`
	RangeKey     string  `json:"range_key"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Body         string  `json:"body"`
	Modified     string  `json:"modified"`
	IsRequired   bool    `json:"is_required"`
	DecisionType *string `json:"type"`
	Status       *string `json:"status"`
}
