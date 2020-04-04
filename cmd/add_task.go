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

		id, _ := cmd.Flags().GetString("id")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		tags, _ := cmd.Flags().GetString("tags")

	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	addTaskCmd.Flags().StringP("id", "i", "", "Id Task")
	addTaskCmd.Flags().StringP("title", "t", "", "Title Task")
	addTaskCmd.PersistentFlags().StringP("description", "d", "", "Description Task")
	addTaskCmd.PersistentFlags().StringP("tags", "g", "", "Tags Task")
}
