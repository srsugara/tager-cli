package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
)

// globalDatatCmd represents the version command
var deleteTaskCmd = &cobra.Command{
	Use:   "delete-task",
	Short: "Delete task from list",
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

		newRev, err := db.Delete(context.TODO(), id, rev)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Task deleted has revision %s\n", newRev)
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
	deleteTaskCmd.PersistentFlags().StringP("id", "i", "", "Id Task")
	deleteTaskCmd.PersistentFlags().StringP("rev", "r", "", "Revision Task")
}
