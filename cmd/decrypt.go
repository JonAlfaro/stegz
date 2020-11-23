package cmd

import (
	"bufio"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/auyer/steganography"
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

		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		if !dirExists(path + "\\.git\\") {
			log.Fatalf("working directory '%s' is not a git repo", path)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		pass := args[0]

		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {

				if err != nil {
					return err
				}

				if strings.HasPrefix(filepath.Base(path), "psr") {

					inFile, _ := os.Open(path) // opening file
					defer inFile.Close()

					reader := bufio.NewReader(inFile)
					img, _ := png.Decode(reader)

					sizeOfMessage := steganography.GetMessageSizeFromImage(img)

					hiddenFile := steganography.Decode(sizeOfMessage, img)

					deFile := decrypt(hiddenFile, pass)
					outFile, _ := os.Create(path) // create file

					outFile.Write(deFile)
					outFile.Close()
				}

				return nil
			})

		if err != nil {
			log.Println(err)
		}
	},
}
