package main

import (
	"log"

	"github.com/martonorova/kubedepend-backend/config"
	"github.com/martonorova/kubedepend-backend/rest"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// Load configuration
	var cfg config.Config
	// TODO read this from env var or commandline parameter
	if err := cfg.LoadConfig("config.yaml"); err != nil {
		log.Panicln(err.Error())
	}

	server, err := rest.NewServer(cfg)
	if err != nil {
		log.Panicln(err.Error())
	}

	server.Start()
}
