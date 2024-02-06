package models

type StageWithStep struct {
	Id           string  `json:"id"`
	PartitionKey string  `json:"partition_key"`
	RangeKey     int     `json:"range_key"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Body         string  `json:"body"`
	Modified     string  `json:"modified"`
	IsRequired   bool    `json:"is_required"`
	Type         *string `json:"type"`
	Status       string  `json:"status"`
	Steps        []*Step `json:"steps"`
}
