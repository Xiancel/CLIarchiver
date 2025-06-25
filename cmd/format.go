package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "List supported archive formats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Supported archive formats:")
		fmt.Println("1. ZIP - Standard zip archive")
		fmt.Println("2. TAR - Uncompressed tar archive")
		fmt.Println("3. TAR.GZ - Compressed tar archive with gzip")
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
