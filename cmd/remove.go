/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove command called",
	Long:  `Remove command called`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {

	rootCmd.AddCommand(removeCmd)
}
