package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/application"
	"github.com/martonorova/kubedepend-backend/dto"
	m "github.com/martonorova/kubedepend-backend/model"
)

func GetJobs(c *gin.Context) {

	var jobs []m.Job

	dbClient := c.MustGet("app").(*application.Application).DB.Client

	result := dbClient.Find(&jobs)

	if result.Error != nil {
		panic("DB error")
	}

	c.IndentedJSON(http.StatusOK, m.APISuccessWithData(jobs))
}

func GetJob(c *gin.Context) {
	dbClient := c.MustGet("app").(*application.Application).DB.Client

	var job m.Job

	if err := dbClient.First(&job, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, m.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, m.APISuccessWithData(job))
}

func AddJob(c *gin.Context) {
	var input dto.CreateJobDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, m.APIError())
		return
	}

	// Save to db
	dbClient := c.MustGet("app").(*application.Application).DB.Client

	job := m.Job{Input: input.Input}
	result := dbClient.Create(&job)

	if result.Error != nil {
		panic("DB error")
	}

	c.IndentedJSON(http.StatusCreated, m.APISuccessWithData(job))
}
