package job

import (
	"fmt"

	"github.com/martonorova/kubedepend-backend/pkg/models"
	"github.com/martonorova/kubedepend-backend/pkg/services"
	"github.com/martonorova/kubedepend-backend/pkg/storage"
)

type defaultJobService struct {
	repository storage.JobRepository
}

func NewDefaultJobService(repository storage.JobRepository) services.JobService {
	return &defaultJobService{
		repository: repository,
	}
}

// defaultJobService implements JobService defined in services.go

func (s *defaultJobService) Create(job models.Job) (*models.Job, error) {
	return s.repository.Create(job)
}

func (s *defaultJobService) FindAll() (*[]models.Job, error) {
	fmt.Printf("%+v\n", s.repository)
	return s.repository.FindAll()
}

func (s *defaultJobService) FindByID(jobID uint64) (*models.Job, error) {
	return s.repository.FindByID(jobID)
}

func (s *defaultJobService) Delete(jobID uint64) error {
	return s.repository.Delete(jobID)
}

func (s *defaultJobService) Update(job models.Job) (*models.Job, error) {
	return s.repository.Update(job)
}
