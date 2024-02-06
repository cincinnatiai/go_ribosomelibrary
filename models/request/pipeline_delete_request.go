package request

type PipelineDeleteRequest struct {
	ClientId     string `json:"client_id"`
	ClientKey    string `json:"client_key"`
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
	IsHardDelete bool   `json:"is_hard_delete"`
}
