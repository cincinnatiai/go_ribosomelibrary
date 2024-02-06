package request

type FetchRequest struct {
	ClientId     string  `json:"client_id"`
	LastRangeKey *string `json:"last_range_key"`
}
