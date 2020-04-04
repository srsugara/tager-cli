package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "tager",
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
