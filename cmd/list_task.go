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
		query := map[string]interface{}{
			"selector": map[string]interface{}{"tags": "frontend"},
			"fields":   []string{"_id", "_rev", "title", "description", "status", "tags"},
		}
		rows, err := db.Find(context.TODO(), query)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var doc interface{}
			if err := rows.ScanDoc(&doc); err != nil {
				panic(err)
			}
			/* do something with doc */
			dat, _ := doc.(map[string]interface{})

		}
	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
