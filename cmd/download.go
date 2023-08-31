/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Jay179-sudo/P2PComet/internal/torrent"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Files",
	Long:  `This command is used to download download and store a torrent file`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		inPath := args[0]
		outPath := args[1]

		tf, err := torrent.Open(inPath)
		if err != nil {
			log.Fatal(err)
		}

		err = tf.DownloadToFile(outPath)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
