package cmd

import (
	"context"

	"github.com/fatih/color"
	"github.com/go-kivik/kivik"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
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
			"selector": map[string]interface{}{},
			"fields":   []string{"_id", "_rev", "title", "description", "status", "tags", "created_date"},
		}
		rows, err := db.Find(context.TODO(), query)
		if err != nil {
			panic(err)
		}
		// instantiate and styling table to show data
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("ID", "Rev", "Title", "Description", "Status", "Tags", "Created")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for rows.Next() {
			var doc interface{}
			if err := rows.ScanDoc(&doc); err != nil {
				panic(err)
			}
			/* do something with doc */
			dat, _ := doc.(map[string]interface{})
			tbl.AddRow(dat["_id"], dat["_rev"], dat["title"], dat["description"], dat["status"], dat["tags"], dat["created_date"])
		}
		tbl.Print()
		if rows.Err() != nil {
			panic(rows.Err())
		}
	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
