package scheduler

import (
	"time"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

type cronScheduler struct {
	*gocron.Scheduler
}

type payload struct {
	TaskName string `json:"taskName"`
	CronJob string `json:"cronJob"`
	//TODO container stuffs
}

func addCronJobs(s *cronScheduler ,cron , taskName string, TaskFunc any) (*gocron.Job, error) {
	job , err := s.Cron(cron).Tag(taskName).Do(TaskFunc)
	if err != nil {
		return nil, err
		}
	s.StartAsync()
	return job, err
}

func listCronJobs(s *cronScheduler) map[string]map[string]any {
	jobs := make(map[string]map[string]any)
	for _, job := range s.Jobs() {
		jobs[job.Tags()[0]] = make(map[string]any)
		jobs[job.Tags()[0]]["isRunning"] = job.IsRunning()
		jobs[job.Tags()[0]]["lastRun"] = job.LastRun()
		jobs[job.Tags()[0]]["nextRun"] = job.NextRun()
		jobs[job.Tags()[0]]["runCount"] = job.RunCount()
		jobs[job.Tags()[0]]["scheduledTime"] = job.ScheduledTime()
	}
	return jobs
}

func deleteCronJob(s *cronScheduler, tags ...string) error {
	return s.RemoveByTags(tags...)
}

func NewCron() *cronScheduler {
	gs:= &cronScheduler{
		gocron.NewScheduler(time.Local),
	}
	gs.TagsUnique()
	return gs
}

func (s *cronScheduler)LetJobs(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": listCronJobs(s),
	})
}

func (s *cronScheduler)AddJob(c *gin.Context) {
	body := payload{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "Invalid inputs, please check your inputs",
			})
		return
	}

	job, err := addCronJobs(s, body.CronJob, body.TaskName, func () {fmt.Println("funciona!")}) //TODO function handler
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"isRunningNow": job.IsRunning(),
		"name": job.Tags(),
		"next run": job.NextRun(),
	})
	
}

func (s *cronScheduler)DeleteJob(c *gin.Context) {
	jobName := c.Param("jobname")

	deleteCronJob(s, jobName)
	c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"result": listCronJobs(s),
	})

}
