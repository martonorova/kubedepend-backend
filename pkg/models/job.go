package models

type JobStatus string

const (
	JOB_CREATED    JobStatus = "CREATED"
	JOB_SUBMITTED  JobStatus = "SUBMITTED"
	JOB_INPROGRESS JobStatus = "INPROGRESS"
	JOB_DONE       JobStatus = "DONE"
	JOB_FAILED     JobStatus = "FAILED"
)

type Job struct {
	ID     uint64    `json:"id" gorm:"primaryKey"`
	Input  uint64    `json:"input"`
	Status JobStatus `json:"status"`
	Result int64     `json:"result"`
}
