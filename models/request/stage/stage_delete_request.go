package stage

type DeleteRequest struct {
	ClientId     string `json:"client_id"`
	ClientKey    string `json:"client_key"`
	PartitionKey string `json:"partition_key"`
	RangeKey     int    `json:"range_key"`
	IsHardDelete bool   `json:"is_hard_delete"`
}
