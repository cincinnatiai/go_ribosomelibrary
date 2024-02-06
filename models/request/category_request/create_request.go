package request

type CreateRequest struct {
	ClientId    string `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
