package application

import (
	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/config"
	c "github.com/martonorova/kubedepend-backend/constants"
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

func (app *Application) StartAPI() error {
	router := gin.Default()
	router.GET(c.ROUTE_ALL_JOB)
	router.POST(c.ROUTE_ALL_JOB)

	err := router.Run("0.0.0.0:8080")

	return err
}
