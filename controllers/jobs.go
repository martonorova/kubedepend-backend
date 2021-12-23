package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/application"
	m "github.com/martonorova/kubedepend-backend/model"
)

func GetJobs(c *gin.Context) {

	var jobs []m.Job

	// result := application.Application.DB.Client.Find(&jobs)
	dbClient := c.MustGet("app").(*application.Application).DB.Client

	result := dbClient.Find(&jobs)

	if result.Error != nil {
		panic("DB error")
	}

	c.IndentedJSON(http.StatusOK, jobs)
}
