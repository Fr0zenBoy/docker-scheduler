package routes

import (
	"net/http"
	"time"
	"context"


	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/go-co-op/gocron"
	"github.com/gin-gonic/gin"

	"github.com/Fr0zenBoy/docker-scheduler/pkg/cron"
	"github.com/Fr0zenBoy/docker-scheduler/pkg/docker"
)

type Job struct {
	Cron string `json:"cron"`
	TaskName string `json:"taskName"`
	ContainerName string `json:"containerName"`
	ContainerConfig container.Config `json:"containerConfig"`
}

type CronDocker struct {
	cli *client.Client
	s   *gocron.Scheduler
}

func (cd *CronDocker) LetJobs(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": cron.ListCronJobs(cd.s),
	})
}

func (cd *CronDocker) AddJob(c *gin.Context) {
	body := Job{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "Invalid inputs, please check your inputs",
			})
		return
	}

	job, err := cron.AddCronJobs(cd.s, body.Cron, body.TaskName, docker.RunContainer, context.Background(), cd.cli, body.ContainerConfig, body.ContainerName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
			"mesage": "fail to create a new cron job",
		})
		return
	}
	
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"isRunningNow": job.IsRunning(),
		"name": job.Tags(),
		"nextRun": job.NextRun(),
	})
	
}

func (cd *CronDocker) DeleteJob(c *gin.Context) {
	jobName := c.Param("jobname")

	cron.DeleteCronJob(cd.s, jobName)
	c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"result": cron.ListCronJobs(cd.s),
	})

}

func NewCronDocker() *CronDocker {
	s:= gocron.NewScheduler(time.Local)
	s.TagsUnique()

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	cronDocker := &CronDocker{
		cli: cli,
		s: s,
	}
	return  cronDocker
}
