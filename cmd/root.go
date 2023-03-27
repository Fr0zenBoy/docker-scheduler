package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "docker-scheduler",
	Short:   "TODO",
	Long:    "TODO",
	Example: "TODO",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "scheduler logs:", err)
		os.Exit(1)
	}
}
