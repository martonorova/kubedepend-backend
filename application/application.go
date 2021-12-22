package application

import (
	"github.com/martonorova/kubedepend-backend/config"
	"github.com/martonorova/kubedepend-backend/db"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnString())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}
