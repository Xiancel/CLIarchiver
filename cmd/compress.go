package cmd

import (
	"cliarchiver/config"

	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compressing file to archive",
	Run: func(cmd *cobra.Command, args []string) {
		config.Input()
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
}
