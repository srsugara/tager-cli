package cmd

import (
	"context"

	"github.com/go-kivik/kivik"
	"github.com/spf13/cobra"
)

// globalDatatCmd represents the version command
var listTaskCmd = &cobra.Command{
	Use:   "list-task",
	Short: "Showing existing task list",
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
	rootCmd.AddCommand(listTaskCmd)
}
