package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

func NewCron() *gocron.Scheduler {
	s := gocron.NewScheduler(time.Local)
	s.TagsUnique()
	return s 
}

func AddCronJobs(s *gocron.Scheduler ,cron, taskName string, TaskFunc any) (*gocron.Job, error) {
	job , err := s.Cron(cron).Tag(taskName).Do(TaskFunc)
	if err != nil {
		return nil, err
		}
	s.StartAsync()
	return job, err
}

func GetCronJobs(s *gocron.Scheduler) []string {
	var jobs []string
	for _, tags := range s.Jobs() {
		
		return tags.Tags()
	}
	return jobs
}

func DeleteCronJob(s *gocron.Scheduler, tags ...string) error {
	return s.RemoveByTags(tags...)
}
