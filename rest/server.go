package rest

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/config"
	"github.com/martonorova/kubedepend-backend/pkg/application"
	"github.com/martonorova/kubedepend-backend/pkg/ports/http/controllers"
	"github.com/martonorova/kubedepend-backend/rest/constants"
)

type Server interface {
	// Close() error
	Start() error
}

type serverControllers struct {
	job controllers.HTTPJobController
}

type defaultServer struct {
	engine      *gin.Engine
	address     string
	controllers serverControllers
}

func NewServer(cfg config.Config) (Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	app, err := application.Get(cfg)
	if err != nil {
		log.Panicln(err.Error())
	}

	app.StartServices()

	jobController := app.JobController

	serverControllers := serverControllers{
		job: jobController,
	}

	server := defaultServer{
		engine: engine,
		// listen on all interfaces
		address:     fmt.Sprintf("0.0.0.0:%d", cfg.Server.Port),
		controllers: serverControllers,
	}

	server.configure(cfg)

	return server, nil
}

func (s defaultServer) Start() error {
	return s.engine.Run(s.address)
}

func (s *defaultServer) configure(cfg config.Config) {
	apiV1Router := s.engine.Group("/api/v1")
	{
		apiV1Router.GET(constants.ROUTE_ALL_JOB, s.controllers.job.GetJobs)
		apiV1Router.GET(constants.ROUTE_SINGLE_JOB, s.controllers.job.GetJob)
		apiV1Router.POST(constants.ROUTE_ALL_JOB, s.controllers.job.AddJob)
		apiV1Router.DELETE(constants.ROUTE_SINGLE_JOB, s.controllers.job.DeleteJob)
	}
}
