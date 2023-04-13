package cron

import (
	"github.com/go-co-op/gocron"
)

func AddCronJobs(s *gocron.Scheduler, cron, taskName string, jobFun any, params ...any) (*gocron.Job, error) {
	job , err := s.Cron(cron).Tag(taskName).Do(jobFun, params...)
	if err != nil {
		return nil, err
		}
	s.StartAsync()
	return job, err
}

func ListCronJobs(s *gocron.Scheduler) map[string]map[string]any {
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

func DeleteCronJob(s *gocron.Scheduler, tags ...string) error {
	return s.RemoveByTags(tags...)
}

