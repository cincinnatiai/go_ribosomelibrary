package models

type Category struct {
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Created      string `json:"created"`
	Modified     string `json:"modified"`
	Status       string `json:"status"`
}
