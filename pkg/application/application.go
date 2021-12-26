package application

import (
	"github.com/martonorova/kubedepend-backend/config"
	"github.com/martonorova/kubedepend-backend/db"
	"github.com/martonorova/kubedepend-backend/services/worker"
)

// struct to hold configuration and db connection
type Application struct {
	DB         *db.DB
	Cfg        *config.Config
	Dispatcher *worker.Dispatcher
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnString())

	if err != nil {
		return nil, err
	}

	if err := db.SetupModels(); err != nil {
		return nil, err
	}

	dispatcher := worker.NewDispatcher(cfg.NWorkers, cfg.JobQueueSize)

	return &Application{
		DB:         db,
		Cfg:        cfg,
		Dispatcher: dispatcher,
	}, nil
}
