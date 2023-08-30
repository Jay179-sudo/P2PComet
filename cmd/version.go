/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version command called",
	Long:  `Version command called`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
	},
}

func init() {

	rootCmd.AddCommand(versionCmd)
}
