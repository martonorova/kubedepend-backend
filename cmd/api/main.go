package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/martonorova/kubedepend-backend/application"
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

	if err := app.StartAPI(); err != nil {
		log.Fatal(err.Error())
	}

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
