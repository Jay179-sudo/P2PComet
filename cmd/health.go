/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Jay179-sudo/P2PComet/internal/torrent"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Checking the health of the torrent file input",
	Long:  `Checking the health of the torrent file input`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		torrentInput := args[0]
		_, err := torrent.Open(torrentInput)
		if err != nil {
			log.Fatal("The Torrent file you have provided is not valid. Please retry with a different file!")
		}
		fmt.Print("The file seems to be OK! To download the contents, use the download command!")
	},
}

func init() {

	rootCmd.AddCommand(healthCmd)
}
