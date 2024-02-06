package campaign

type FetchAllRequest struct {
	ClientId             string  `json:"client_id"`
	ClientKey            string  `json:"client_key"`
	PipelinePartitionKey string  `json:"pipeline_partition_key"`
	PipelineRangeKey     string  `json:"pipeline_range_key"`
	LastRangeKey         *string `json:"last_range_key"`
}
