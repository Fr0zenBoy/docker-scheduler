package scheduler

import (
	"io"
	"os"
	"time"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/go-co-op/gocron"
)

func CreateTask(cron string, f io.ReadCloser) {
	//exemple "*/1 * * * *"
	s := gocron.NewScheduler(time.Local)
	s.Cron(cron).Do(f)
	s.StartAsync()

	stdcopy.StdCopy(os.Stdout, os.Stderr, f)
}
