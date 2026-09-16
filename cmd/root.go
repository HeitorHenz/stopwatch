package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stopwatch",
	Short: "A simple stopwatch CLI",
	Long:  `A simple stopwatch CLI that can be used to measure elapsed time.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("stopwatch")
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
