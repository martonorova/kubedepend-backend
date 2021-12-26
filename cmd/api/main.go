package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	c "github.com/martonorova/kubedepend-backend/constants"
	"github.com/martonorova/kubedepend-backend/controllers"
	"github.com/martonorova/kubedepend-backend/exithandler"
	"github.com/martonorova/kubedepend-backend/pkg/application"
	"github.com/martonorova/kubedepend-backend/services/worker"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars from .env file")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	collector := worker.NewCollector(100)

	app.Dispatcher.Start(collector)

	if err := startAPI(app); err != nil {
		log.Fatal(err.Error())
	}

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}

func startAPI(app *application.Application) error {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.Use(func(c *gin.Context) {
		// provide the application instance to controllers
		// TODO only provide this to selected controllers
		c.Set("app", app)
		c.Next()
	})

	router.GET(c.ROUTE_ALL_JOB, controllers.GetJobs)
	router.GET(c.ROUTE_SINGLE_JOB, controllers.GetJob)
	router.POST(c.ROUTE_ALL_JOB, controllers.AddJob)
	router.DELETE(c.ROUTE_SINGLE_JOB, controllers.DeleteJob)

	err := router.Run("0.0.0.0:8080")

	return err
}
