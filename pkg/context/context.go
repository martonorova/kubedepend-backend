package context

import (
	"errors"

	"github.com/martonorova/kubedepend-backend/config"
	"github.com/martonorova/kubedepend-backend/pkg/ports/http/controllers"
	"github.com/martonorova/kubedepend-backend/pkg/services"
	"github.com/martonorova/kubedepend-backend/pkg/services/collector"
	"github.com/martonorova/kubedepend-backend/pkg/services/execution"
	"github.com/martonorova/kubedepend-backend/pkg/services/job"
	"github.com/martonorova/kubedepend-backend/pkg/storage"
	"github.com/martonorova/kubedepend-backend/pkg/storage/sql"
)

// based on the configuration, it returns resources

func NewJobRepository(cfg config.Config) (storage.JobRepository, error) {
	// init storage connection based on config

	if cfg.Storage.SQL.Enabled {
		return sql.NewSQLJobRepository(cfg.Storage.SQL.GetDBConnString())
	}

	return nil, errors.New("no storage enabled")
}

func NewJobService(cfg config.Config, jobRepository storage.JobRepository) (services.JobService, error) {

	return job.NewDefaultJobService(
		jobRepository,
	), nil
}

func NewCollectorService(cfg config.Config, jobService services.JobService) (services.CollectorService, error) {
	return collector.NewDefaultCollectorService(jobService), nil
}

func NewExecutionService(cfg config.Config, collectorService services.CollectorService) (services.ExecutionService, error) {

	return execution.NewDefaultExecutionService(
		cfg.Worker.WorkerCount,
		cfg.Worker.QueueSize,
		collectorService,
	), nil
}

func NewHTTPJobController(
	cfg config.Config,
	jobService services.JobService,
	executionService services.ExecutionService,
) (controllers.HTTPJobController, error) {

	return controllers.NewHTTPJobController(
		jobService,
		executionService,
	), nil
}
