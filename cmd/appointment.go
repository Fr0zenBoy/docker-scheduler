package cmd

import (
	"github.com/Fr0zenBoy/docker-scheduler/container"
	"github.com/Fr0zenBoy/docker-scheduler/scheduler"
	"github.com/spf13/cobra"
)

var appointmentCmd = &cobra.Command{
	Use:     "appointment",
	Aliases: []string{"appoit"},
	Short:   "TODO",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		scheduler.CreateTask(args[1], container.RunContainer(args[0])) 
	},
}

func init() {
	rootCmd.AddCommand(appointmentCmd)
}
