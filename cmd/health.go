/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Checking the health of the torrent file input",
	Long:  `Checking the health of the torrent file input`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("health called")
	},
}

func init() {

	rootCmd.AddCommand(healthCmd)
}
