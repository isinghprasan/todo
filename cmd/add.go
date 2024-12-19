/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/isinghprasan/todo/td"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item in the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items := []td.Item{}

	for _, x := range args {
		item := td.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err := td.SaveItems(datafile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
