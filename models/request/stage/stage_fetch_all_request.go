package stage

type StageFetchAllRequest struct {
	ClientId   string `json:"client_id"`
	ClientKey  string `json:"client_key"`
	PipelineId string `json:"pipeline_id"`
}
