package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCMD)
}

var encryptCMD = &cobra.Command{
	Use:   "e [ENCRYPTION KEY]",
	Short: "encrypt all files recursively from the current working directory that have the prefix 'psr'",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Expected 1 arg for [ENCRYPTION KEY], but recieved '%v'", len(args))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}
