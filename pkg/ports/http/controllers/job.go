package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martonorova/kubedepend-backend/pkg/models"
	"github.com/martonorova/kubedepend-backend/pkg/ports/http/requests"
	"github.com/martonorova/kubedepend-backend/pkg/ports/http/responses"
	"github.com/martonorova/kubedepend-backend/pkg/services"
)

type HTTPJobController struct {
	jobService       services.JobService
	executionService services.ExecutionService
}

func NewHTTPJobController(jobService services.JobService, executionService services.ExecutionService) HTTPJobController {
	return HTTPJobController{
		jobService:       jobService,
		executionService: executionService,
	}
}

func (jc HTTPJobController) GetJobs(c *gin.Context) {
	jobs, err := jc.jobService.FindAll()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, responses.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, responses.APISuccessWithData(jobs))
}

func (jc HTTPJobController) GetJob(c *gin.Context) {

	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, responses.APIError())
		return
	}

	job, err := jc.jobService.FindByID(jobID)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotFound, responses.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, responses.APISuccessWithData(job))
}

func (jc HTTPJobController) AddJob(c *gin.Context) {
	var input requests.CreateJobRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.APIError())
		return
	}

	// Save Job
	job, err := jc.jobService.Create(models.Job{Input: input.Input, Status: models.JobStatusCreated})
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, responses.APIError())
		return
	}

	// TODO
	// dispatcher := c.MustGet("app").(*application.Application).Dispatcher

	// // Collect Job for worker pool
	// if err := dispatcher.Submit(dto.SubmitJobDTO{ID: job.ID, Input: job.Input}); err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, m.APIError())
	// }

	jc.executionService.SubmitJob(&services.SubmitJobDTO{ID: job.ID, Input: job.Input})

	c.IndentedJSON(http.StatusCreated, responses.APISuccessWithData(job))
}

func (jc HTTPJobController) DeleteJob(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, responses.APIError())
		return
	}

	if err := jc.jobService.Delete(jobID); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusNotFound, responses.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, responses.APISuccess())
}
