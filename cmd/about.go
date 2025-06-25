package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Display information about the Archiver program",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("- Archiver v1.0")
		fmt.Println("- Simple CLI file archiver")
		fmt.Println("- Supports formats: zip, tar, tar.gz")
		fmt.Println("- Open source project")
		fmt.Println("- Repository: https://github.com/Xiancel/CLIarchiver.git")
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
