/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List does something",
	Long:  `List does something`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
