package request

type StepInputFetchByUserRequest struct {
	ClientId  string `json:"client_id"`
	ClientKey string `json:"client_key"`
	UserId    string `json:"user_id"`
}
