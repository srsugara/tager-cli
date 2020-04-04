package cmd

import (
	"context"

	"github.com/spf13/cobra"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
)

// globalDatatCmd represents the version command
var addTaskCmd = &cobra.Command{
	Use:   "add-task",
	Short: "Add task to list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := kivik.New("couch", "http://13.250.43.79:5984/")
		if err != nil {
			panic(err)
		}

		db := client.DB(context.TODO(), "tager")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}
