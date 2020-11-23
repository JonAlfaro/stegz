package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		if !dirExists(path + "\\.git\\") {
			log.Fatalf("working directory '%s' is not a git repo", path)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {

				if err != nil {
					return err
				}

				fmt.Println(path, info.Size())
				return nil
			})

		if err != nil {
			log.Println(err)
		}
	},
}
