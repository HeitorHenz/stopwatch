package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"stopwatch/internal/ui"
)

var rootCmd = &cobra.Command{
	Use:          "stopwatch",
	Short:        "A simple stopwatch CLI",
	Long:         `A simple stopwatch CLI that can be used to measure elapsed time.`,
	Args:         cobra.NoArgs,
	Version:      "0.1.0",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(ui.New(), tea.WithAltScreen())
		_, err := p.Run()
		return err
	},
}

func Execute() error {
	return rootCmd.Execute()
}
