package cmd

import (
	"github.com/spf13/cobra"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
)

// globalDatatCmd represents the version command
var updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Update task from list",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	updateTaskCmd.PersistentFlags().StringP("id", "i", "", "Id Task")
	updateTaskCmd.PersistentFlags().StringP("rev", "r", "", "Revision Task")
	updateTaskCmd.PersistentFlags().StringP("title", "t", "", "Title Task")
	updateTaskCmd.PersistentFlags().StringP("description", "d", "", "Description Task")
	updateTaskCmd.PersistentFlags().StringP("status", "s", "", "Status Task")
	updateTaskCmd.PersistentFlags().StringP("tags", "g", "", "Tags Task")
}
