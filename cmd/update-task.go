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
var updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Update task from list",
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
		rev, _ := cmd.Flags().GetString("rev")
		status, _ := cmd.Flags().GetString("status")

		doc := Task{
			Rev:     rev,
			Status:  status,
			DirtyAt: "" + time.Now().Format(time.RFC3339),
		}
		newRev, err := db.Put(context.TODO(), id, doc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Task updated has new revision%s\n", newRev)
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
	updateTaskCmd.PersistentFlags().StringP("id", "i", "", "Id Task")
	updateTaskCmd.PersistentFlags().StringP("rev", "r", "", "Revision Task")
	updateTaskCmd.PersistentFlags().StringP("status", "s", "", "Status Task")
}
