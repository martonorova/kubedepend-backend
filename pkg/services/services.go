package services

import (
	"github.com/martonorova/kubedepend-backend/pkg/models"
)

// the ports/adapters/controllers should use this to interact with the domain layer
type JobService interface {
	Create(job models.Job) (*models.Job, error)
	FindAll() (*[]models.Job, error)
	FindByID(jobID uint64) (*models.Job, error)
	Delete(jobID uint64) error
	Update(job models.Job) (*models.Job, error)
}

type ExecutionService interface {
	SubmitJob(jobSubmit *SubmitJobDTO) error
	Start()
}

type CollectorService interface {
	CollectJob(jobResult *JobResultDTO)
	Start()
}

type SubmitJobDTO struct {
	ID    uint64
	Input uint64
}

type JobResultDTO struct {
	ID     uint64
	Result uint64
}
