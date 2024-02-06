package request

type StepFetchRequest struct {
	ClientId   string `json:"client_id"`
	ClientKey  string `json:"client_key"`
	PipelineId string `json:"pipeline_id"`
	StageId    string `json:"stage_id"`
}
