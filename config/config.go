package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	dbUser       string
	dbPassword   string
	dbHost       string
	dbPort       string
	dbName       string
	NWorkers     uint64
	JobQueueSize uint64
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&conf.dbPassword, "dbpassword", os.Getenv("POSTGRES_PASSWORD"), "DB user password")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "DB host")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "DB port")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "DB name")

	nWorkers, err := strconv.ParseUint(os.Getenv("WORKER_COUNT"), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	jobQueueSize, err := strconv.ParseUint(os.Getenv("QUEUE_SIZE"), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	flag.Uint64Var(&conf.NWorkers, "nworkers", nWorkers, "Number of workers")
	flag.Uint64Var(&conf.JobQueueSize, "queuesize", jobQueueSize, "Size of job queue")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPassword,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}
