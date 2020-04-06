package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          string `json:"_id,omitempty"`
	Rev         string `json:"_rev,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Tags        string `json:"tags,omitempty"`
	Status      string `json:"status,omitempty"`
	DirtyAt     string `json:"dirtyAt"`
}

var rootCmd = &cobra.Command{
	Use: "tager-cli",
	Short: `
                        
 /__ __\/  _ \/  __//  __//  __\      /   _\/ \   / \
 / \  | / \|| |  _|  \  |  \/|_____ |  /  | |   | |
 | |  | |-||| |_//|  /_ |    /\____\|  \__| |_/\| |
 \_/  \_/ \|\____\\____\\_/\_\      \____/\____/\_
 ____________________________________________________
                             
 Author      : Syauqi Rahmat Sugara                                                                   
 Email       : syauqisugara@gmail.com                         
 Data Source : Efishery CouchDB`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
