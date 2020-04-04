package cmd

import (
	"github.com/spf13/cobra"
)

// globalDatatCmd represents the version command
var listTaskCmd = &cobra.Command{
	Use:   "list-task",
	Short: "Showing existing task list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
