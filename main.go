package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Fr0zenBoy/docker-scheduler/scheduler"
	"github.com/Fr0zenBoy/docker-scheduler/container"
	fvd "github.com/docker/docker/api/types/container"
)

var task = func() {}

func main() {

	container.NewContainer("nginx", "test", fvd.Config{Tty: false}).RunContainer()

	s := scheduler.NewCron()

	app := gin.New()

	app.GET("/api/jobs", s.LetJobs)
	app.POST("/api/jobs", s.AddJob)
	app.DELETE("/api/jobs/:jobname", s.DeleteJob)

  if err := app.Run(":9092"); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
