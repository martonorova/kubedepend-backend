package dto

type CreateJobRequest struct {
	Input uint64 `json:"input"`
}

type SubmitJobDTO struct {
	ID    uint64
	Input uint64
}

type JobResultDTO struct {
	ID     uint64
	Result uint64
}
