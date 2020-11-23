package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCMD)
}

var decryptCMD = &cobra.Command{
	Use:   "d [ENCRYPTION KEY]",
	Short: "decrypt all files recursively from the current working directory that have the prefix 'psr'",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Expected 1 arg for [ENCRYPTION KEY], but recieved '%v'", len(args))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}
