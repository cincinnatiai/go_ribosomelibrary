package request

type PipelineCreateRequest struct {
	ClientId             string  `json:"client_id"`
	ClientKey            string  `json:"client_key"`
	IdentifyingCategory  string  `json:"identifying_category"`
	Title                string  `json:"title"`
	Description          string  `json:"description"`
	IsPublic             bool    `json:"is_public"`
	AuxiliarPartitionKey *string `json:"auxiliar_partition_key"`
	AuxiliarRangeKey     *string `json:"auxiliar_range_key"`
	Type                 *string `json:"type"`
}
