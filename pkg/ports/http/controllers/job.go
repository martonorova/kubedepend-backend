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

type JobController struct {
	jobService services.JobService
}

func NewJobController(jobService services.JobService) *JobController {
	return &JobController{
		jobService: jobService,
	}
}

func (jc *JobController) GetJobs(c *gin.Context) {
	jobs, err := jc.jobService.FindAll()
	if err != nil {
		log.Panicln(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, responses.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, responses.APISuccessWithData(jobs))
}

func (jc *JobController) GetJob(c *gin.Context) {

	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Panicln(err.Error())
		c.IndentedJSON(http.StatusBadRequest, responses.APIError())
		return
	}

	job, err := jc.jobService.FindByID(jobID)
	if err != nil {
		log.Panicln(err.Error())
		c.IndentedJSON(http.StatusNotFound, responses.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, responses.APISuccessWithData(job))
}

func (jc *JobController) AddJob(c *gin.Context) {
	var input requests.CreateJobRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.APIError())
		return
	}

	// Save Job
	job, err := jc.jobService.Create(models.Job{Input: input.Input})
	if err != nil {
		log.Panicln(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, responses.APIError())
		return
	}

	// TODO
	// dispatcher := c.MustGet("app").(*application.Application).Dispatcher

	// // Collect Job for worker pool
	// if err := dispatcher.Submit(dto.SubmitJobDTO{ID: job.ID, Input: job.Input}); err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, m.APIError())
	// }

	c.IndentedJSON(http.StatusCreated, responses.APISuccessWithData(job))
}

func (jc *JobController) DeleteJob(c *gin.Context) {
	jobID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Panicln(err.Error())
		c.IndentedJSON(http.StatusBadRequest, responses.APIError())
		return
	}

	if err := jc.jobService.Delete(jobID); err != nil {
		log.Panicln(err.Error())
		c.IndentedJSON(http.StatusNotFound, responses.APIError())
		return
	}

	c.IndentedJSON(http.StatusOK, responses.APISuccess())
}
