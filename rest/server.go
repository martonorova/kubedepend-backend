package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/pkg/context"
	"github.com/martonorova/kubedepend-backend/pkg/ports/http/controllers"
)

type Server interface {
	// Close() error
	Start(address string) error
}

type serverControllers struct {
	job controllers.JobController
}

type defaultServer struct {
	engine      *gin.Engine
	controllers serverControllers
}

// TODO pass configuration here
func NewServer() (Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	// TODO pass configuration
	jobController, err := context.NewJobController()
	if err != nil {
		log.Panicln(err.Error())
	}

	serverControllers := serverControllers{
		job: jobController,
	}

	return defaultServer{
		engine:      engine,
		controllers: serverControllers,
	}, nil
}

func (s defaultServer) Start(address string) error {
	return s.engine.Run(address)
}
