package campaign

type CreateRequest struct {
	ClientId             string `json:"client_id"`
	ClientKey            string `json:"client_key"`
	PipelinePartitionKey string `json:"pipeline_partition_key"`
	PipelineRangeKey     string `json:"pipeline_range_key"`
	Title                string `json:"title"`
	Description          string `json:"description"`
	SecondaryId          string `json:"secondary_id"`
	CreatorUserId        string `json:"creator_user_id"`
}
