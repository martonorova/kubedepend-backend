package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/martonorova/kubedepend-backend/application"
	c "github.com/martonorova/kubedepend-backend/constants"
	"github.com/martonorova/kubedepend-backend/controllers"
	"github.com/martonorova/kubedepend-backend/exithandler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars from .env file")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

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
		c.Set("app", app)
		c.Next()
	})

	router.GET(c.ROUTE_ALL_JOB, controllers.GetJobs)
	router.POST(c.ROUTE_ALL_JOB)

	err := router.Run("0.0.0.0:8080")

	return err
}
