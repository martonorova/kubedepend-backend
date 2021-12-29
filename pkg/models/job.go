package models

type JobStatus string

const (
	JobStatusCreated    JobStatus = "CREATED"
	JobStatusSubmitted  JobStatus = "SUBMITTED"
	JobStatusInProgress JobStatus = "INPROGRESS"
	JobStatusDone       JobStatus = "DONE"
	JobStatusFailed     JobStatus = "FAILED"
)

type Job struct {
	ID     uint64    `json:"id" gorm:"primaryKey"`
	Input  uint64    `json:"input"`
	Status JobStatus `json:"status"`
	Result int64     `json:"result"`
}
