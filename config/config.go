package config

import (
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
	Worker  WorkerConfig  `yaml:"worker"`
}

type ServerConfig struct {
	Enabled bool   `yaml:"enabled", envconfig:"SERVER_ENABLED"`
	Port    uint64 `yaml:"port", envconfig:"SERVER_PORT"`
}

type StorageConfig struct {
	SQL SQLConfig `yaml:"sql"`
}

type SQLConfig struct {
	Enabled    bool   `yaml:"enabled" envconfig:"STORAGE_SQL_ENABLED"`
	DBUser     string `yaml:"user" envconfig:"STORAGE_SQL_DBUSER"`
	DBPassword string `yaml:"password" envconfig:"STORAGE_SQL_DBPASSWORD"`
	DBHost     string `yaml:"host" envconfig:"STORAGE_SQL_DBHOST"`
	DBPort     uint64 `yaml:"port" envconfig:"STORAGE_SQL_DBPORT"`
	DBName     string `yaml:"dbname" envconfig:"STORAGE_SQL_DBNAME"`
}

type WorkerConfig struct {
	Enabled     bool   `yaml:"enabled" envconfig:"WORKER_ENABLED"`
	WorkerCount uint64 `yaml:"workerCount" envconfig:"WORKER_COUNT"`
	QueueSize   uint64 `yaml:"queueSize" envconfig:"QUEUE_SIZE"`
}

func (c *Config) LoadConfig(configFilePath string) error {

	if err := c.readConfigFile(configFilePath); err != nil {
		log.Panicln(err.Error())
		return err
	}

	if err := c.readEnvVars(); err != nil {
		log.Panicln(err.Error())
		return err
	}

	fmt.Printf("%+v\n", c)

	return nil
}

func (c *Config) readConfigFile(configFilePath string) error {
	f, err := os.Open(configFilePath)
	if err != nil {
		log.Panicln(err.Error())
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		log.Panicln(err.Error())
		return err
	}

	return nil
}

func (c *Config) readEnvVars() error {
	err := envconfig.Process("", c)
	if err != nil {
		log.Panicln(err.Error())
		return err
	}

	return nil
}

func (c *SQLConfig) GetDBConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
