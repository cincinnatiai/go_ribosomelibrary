package campaign

type FetchByUser struct {
	ClientId     string  `json:"client_id"`
	ClientKey    string  `json:"client_key"`
	UserId       string  `json:"user_id"`
	LastRangeKey *string `json:"last_range_key"`
}
