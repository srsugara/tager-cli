package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
)

type Task struct {
	ID          string `json:"_id"`
	Rev         string `json:"_rev,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Status      string `json:"status"`
	CreatedDate string `json:"created_date"`
}

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
		doc := Task{
			ID:          id,
			Title:       title,
			Description: description,
			Tags:        tags,
			Status:      "unstarted",
			CreatedDate: "" + time.Now().Format(time.RFC3339),
		}
		rev, err := db.Put(context.TODO(), id, doc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Task inserted has revision%s\n", rev)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	addTaskCmd.PersistentFlags().StringP("id", "i", "", "Id Task")
	addTaskCmd.PersistentFlags().StringP("title", "t", "", "Title Task")
	addTaskCmd.PersistentFlags().StringP("description", "d", "", "Description Task")
	addTaskCmd.PersistentFlags().StringP("tags", "g", "", "Tags Task")
}
