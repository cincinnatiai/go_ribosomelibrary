package stage

type StageCreateRequest struct {
	ClientId       string  `json:"client_id"`
	ClientKey      string  `json:"client_key"`
	PipelineId     string  `json:"pipeline_id"`
	SequenceNumber int     `json:"sequence_number"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Body           string  `json:"body"`
	IsRequired     bool    `json:"is_required"`
	Type           *string `json:"type"`
}
