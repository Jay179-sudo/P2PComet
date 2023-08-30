/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information content",
	Long:  `Information content`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {

	rootCmd.AddCommand(infoCmd)
}
