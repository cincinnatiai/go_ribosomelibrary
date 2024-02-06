package request

type PipelineFetchRequest struct {
	ClientId            string  `json:"client_id"`
	ClientKey           string  `json:"client_key"`
	IdentifyingCategory string  `json:"identifying_category"`
	LastRangeKey        *string `json:"last_range_key"`
}
