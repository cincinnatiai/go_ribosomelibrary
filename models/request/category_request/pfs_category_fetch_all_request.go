package request

type PFSCategoryFetchAllRequest struct {
	ClientId     string  `json:"client_id"`
	ClientKey    string  `json:"client_key"`
	LastRangeKey *string `json:"last_range_key"`
}
