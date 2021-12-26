package sql

import (
	"database/sql"
	"log"

	"github.com/martonorova/kubedepend-backend/pkg/storage"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/martonorova/kubedepend-backend/pkg/models"
	"gorm.io/driver/postgres"
)

type sqlJobRepository struct {
	db gorm.DB
}

func NewSQLJobRepository(connString string) (storage.JobRepository, error) {
	sqlDB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Panic(err.Error())
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	repository := &sqlJobRepository{db: *gormDB}

	return repository, nil
}

//sqlJobRepository implements JobRepository defined in repositories.go

func (r *sqlJobRepository) Create(job models.Job) (*models.Job, error) {

	if err := r.db.Create(&job).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &job, nil
}

func (r *sqlJobRepository) FindAll() (*[]models.Job, error) {
	var jobs []models.Job

	if err := r.db.Find(&jobs).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &jobs, nil
}

func (r *sqlJobRepository) FindByID(jobID uint64) (*models.Job, error) {
	var job models.Job

	if err := r.db.First(&job, jobID).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &job, nil

}
func (r *sqlJobRepository) Delete(jobID uint64) error {
	// Get model if exists
	job, err := r.FindByID(jobID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := r.db.Delete(&job).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (r *sqlJobRepository) Update(job models.Job) (*models.Job, error) {
	// Get model if exists
	oldjob, err := r.FindByID(job.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err := r.db.Model(oldjob).Updates(job).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	updatedJob := oldjob

	return updatedJob, nil

}

func (r *sqlJobRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	sqlDB.Close()
	return nil
}
