package storage

import (
	"github.com/martonorova/kubedepend-backend/pkg/models"
)

// JobRepository for interacting with jobs
type JobRepository interface {
	Create(job models.Job) (*models.Job, error)
	FindAll() (*[]models.Job, error)
	FindByID(jobID uint64) (*models.Job, error)
	Delete(jobID uint64) error
	Update(job models.Job) (*models.Job, error)
	Close() error
}
