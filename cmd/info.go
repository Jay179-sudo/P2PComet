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

var megaByte bool
var gigaByte bool
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information of the input torrent file",
	Long:  `Information content`,
	Run: func(cmd *cobra.Command, args []string) {
		torrentFile, err := torrent.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}

		if torrentFile.Name != "" {
			fmt.Println("Name: ", torrentFile.Name)
		} else {
			fmt.Println("Name: Unknown")
		}

		if megaByte {
			fmt.Println("Torrent Length in Megabytes: ", float64(torrentFile.Length)/(1<<20))
		} else if gigaByte {
			fmt.Println("Torrent Length in Gigabytes: ", float64(torrentFile.Length)/(1<<30))
		} else {
			fmt.Println("Torrent Length in Kilobytes: ", float64(torrentFile.Length)/(1<<10))
		}
		fmt.Println("Piece Length in Kilobytes: ", float64(torrentFile.PieceLength)/(1<<10))
	},
}

func init() {
	infoCmd.Flags().BoolVarP(&megaByte, "megabytes", "m", false, "Display Size in MegaBytes")
	infoCmd.Flags().BoolVarP(&gigaByte, "gigabytes", "g", false, "Display Size in GigaBytes")
	rootCmd.AddCommand(infoCmd)
}
