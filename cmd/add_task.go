package cmd

import (
	"github.com/spf13/cobra"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
)

// globalDatatCmd represents the version command
var addTaskCmd = &cobra.Command{
	Use:   "add-task",
	Short: "Add task to list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}
