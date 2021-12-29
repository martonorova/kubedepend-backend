package application

import (
	"github.com/martonorova/kubedepend-backend/config"
	"github.com/martonorova/kubedepend-backend/pkg/context"
	"github.com/martonorova/kubedepend-backend/pkg/ports/http/controllers"
	"github.com/martonorova/kubedepend-backend/pkg/services"
	"github.com/martonorova/kubedepend-backend/pkg/storage"
)

type Application struct {
	jobRepository    storage.JobRepository
	jobService       services.JobService
	collectorService services.CollectorService
	executionService services.ExecutionService
	JobController    controllers.HTTPJobController
}

func Get(cfg config.Config) (*Application, error) {
	jobRepo, err := context.NewJobRepository(cfg)
	if err != nil {
		return nil, err
	}

	jobSvc, err := context.NewJobService(cfg, jobRepo)
	if err != nil {
		return nil, err
	}

	collectorSvc, err := context.NewCollectorService(cfg, jobSvc)
	if err != nil {
		return nil, err
	}

	executionSvc, err := context.NewExecutionService(cfg, collectorSvc)
	if err != nil {
		return nil, err
	}

	jobController, err := context.NewHTTPJobController(cfg, jobSvc, executionSvc)
	if err != nil {
		return nil, err
	}

	return &Application{
		jobRepository:    jobRepo,
		jobService:       jobSvc,
		collectorService: collectorSvc,
		executionService: executionSvc,
		JobController:    jobController,
	}, nil
}

func (app *Application) StartServices() {
	app.collectorService.Start()
	app.executionService.Start()
}
