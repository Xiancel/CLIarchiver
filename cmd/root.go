package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archiver",
	Short: "Archiver CLI app.",
	Long:  "Archiver is a simple CLI program for archiving files in different formats.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to archiver CLI! Use --help for commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
