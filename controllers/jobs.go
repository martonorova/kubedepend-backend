package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/dto"
	"github.com/martonorova/kubedepend-backend/pkg/application"
	m "github.com/martonorova/kubedepend-backend/pkg/models"
	"gorm.io/gorm"
)

func GetJobs(c *gin.Context) {

	var jobs []m.Job

	dbClient := getDbClientFromContext(c)

	result := dbClient.Find(&jobs)

	if result.Error != nil {
		panic("DB error")
	}

	c.IndentedJSON(http.StatusOK, m.APISuccessWithData(jobs))
}

func GetJob(c *gin.Context) {
	dbClient := getDbClientFromContext(c)

	var job m.Job

	if err := dbClient.First(&job, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, m.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, m.APISuccessWithData(job))
}

func AddJob(c *gin.Context) {
	var input dto.CreateJobRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, m.APIError())
		return
	}

	// Save to db
	dbClient := getDbClientFromContext(c)

	job := m.Job{Input: input.Input, Status: m.JOB_CREATED}
	result := dbClient.Create(&job)

	if result.Error != nil {
		panic("DB error")
	}

	dispatcher := c.MustGet("app").(*application.Application).Dispatcher

	// Collect Job for worker pool
	if err := dispatcher.Submit(dto.SubmitJobDTO{ID: job.ID, Input: job.Input}); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, m.APIError())
	}

	c.IndentedJSON(http.StatusCreated, m.APISuccessWithData(job))
}

func DeleteJob(c *gin.Context) {
	dbClient := getDbClientFromContext(c)

	// Get model if exists
	var job m.Job

	if err := dbClient.First(&job, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, m.APIError())
		return
	}

	result := dbClient.Delete(&job)
	if result.Error != nil {
		panic("DB error")
	}

	c.IndentedJSON(http.StatusOK, m.APISuccess())
}

func getDbClientFromContext(c *gin.Context) *gorm.DB {
	return c.MustGet("app").(*application.Application).DB.Client
}
