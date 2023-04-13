package main

import (
	"os"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Fr0zenBoy/docker-scheduler/pkg/routes"

)

var task = func() {}

func main() {

	app := gin.New()

	cronDocker := routes.NewCronDocker()


	app.GET("/api/jobs", cronDocker.LetJobs)
	app.POST("/api/jobs", cronDocker.AddJob)
	app.DELETE("/api/jobs/:jobname", cronDocker.DeleteJob)

  if err := app.Run(":9092"); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
