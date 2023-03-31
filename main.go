package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/Fr0zenBoy/docker-scheduler/scheduler"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

type payload struct {
	TaskName string `json:"taskName"`
	CronJob string `json:"cronJob"`
	//TODO container stuffs
}

type outGetJobs struct {
	Result []string
}

type ds struct {
 New *gocron.Scheduler
}

func (s ds)GetJobs(c *gin.Context) {
	payload := outGetJobs{
		Result: scheduler.GetCronJobs(s.New),
	}

	c.JSON(http.StatusOK, payload)
}

func (s ds)AddJob(c *gin.Context) {
	body := payload{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "Invalid inputs, please check your inputs",
			})
		return
	}

	job, err := scheduler.AddCronJobs(s.New, body.CronJob, body.TaskName, func () {fmt.Println("funciona!")})
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"isRunning": job.IsRunning(),
		"name": job.Tags(),
		"next run": job.NextRun(),
	})
	
}

func (s ds)DeleteJob(c *gin.Context) {
	jobName := c.Param("jobname")

	scheduler.DeleteCronJob(s.New, jobName)
	c.AbortWithStatus(http.StatusOK)

}

func main() {
	s := ds {
		New: scheduler.NewCron(),
	} 
	
	app := gin.New()

	// api routes
	app.GET("/api/jobs", s.GetJobs)
	app.POST("/api/jobs", s.AddJob)
	app.DELETE("/api/jobs/:id", s.DeleteJob)

	// run api on port 9092
  if err := app.Run(":9092"); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
