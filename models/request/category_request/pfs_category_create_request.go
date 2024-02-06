package request

type PFSCategoryCreateRequest struct {
	ClientId    string `json:"client_id"`
	ClientKey   string `json:"client_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
