package cmd

import (
	"context"
	"fmt"
	"time"

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

		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		tags, _ := cmd.Flags().GetString("tags")
		doc := Task{
			Title:       title,
			Description: description,
			Tags:        tags,
			Status:      "unstarted",
			DirtyAt:     "" + time.Now().Format(time.RFC3339),
		}
		var id, rev string
		id, rev, err = db.CreateDoc(context.TODO(), doc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Task inserted has id:%s and revision:%s\n", id, rev)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	addTaskCmd.PersistentFlags().StringP("title", "t", "", "Title Task")
	addTaskCmd.PersistentFlags().StringP("description", "d", "", "Description Task")
	addTaskCmd.PersistentFlags().StringP("tags", "g", "", "Tags Task")
}
