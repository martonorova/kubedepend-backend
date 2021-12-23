package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/martonorova/kubedepend-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func Get(connString string) (*DB, error) {
	db, err := get(connString)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	sqlDB, err := d.Client.DB()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	sqlDB.Close()
	return nil
}

func (d *DB) SetupModels() error {
	err := d.Client.AutoMigrate(&model.Job{})

	return err
}

func get(connString string) (*gorm.DB, error) {

	sqlDB, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB, nil
}
