package context

import (
	"log"

	"github.com/martonorova/kubedepend-backend/pkg/ports/http/controllers"
	"github.com/martonorova/kubedepend-backend/pkg/services"
	"github.com/martonorova/kubedepend-backend/pkg/services/job"
	"github.com/martonorova/kubedepend-backend/pkg/storage"
)

// based on the configuration, it returns resources

func NewJobRepository() (storage.JobRepository, error) {
	// init storage connection based on config
	return nil, nil
}

func NewJobService() (services.JobService, error) {
	jobRepository, err := NewJobRepository()
	if err != nil {
		log.Panicln(err.Error())
	}

	return job.NewDefaultJobService(
		jobRepository,
	), nil
}

func NewJobController() (controllers.JobController, error) {
	jobService, err := NewJobService()
	if err != nil {
		log.Panicln(err.Error())
	}

	return controllers.NewJobController(
		jobService,
	), nil
}
