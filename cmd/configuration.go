/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "Configuration Mode",
	Long:  `Configuration Mode`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config Mode")
	},
}

func init() {
	rootCmd.AddCommand(confCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
