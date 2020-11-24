package cmd

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alistanis/goenc/aes/gcm"
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
		// pass := args[0]
		dogFile, _ := os.Open("dog.png")      // opening file
		dogReader := bufio.NewReader(dogFile) // buffer reader
		imgDog, _ := png.Decode(dogReader)
		defer dogFile.Close()

		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {

				if err != nil {
					return err
				}

				if strings.HasPrefix(filepath.Base(path), "psr") && strings.HasSuffix(filepath.Base(path), ".png") {
					fmt.Println("hidden file found ", path)

					inFile, _ := os.Open(path) // opening file
					defer inFile.Close()

					reader := bufio.NewReader(inFile)
					img, err := png.Decode(reader)
					if err != nil {
						panic(err)
					}

					if !img.Bounds().Eq(imgDog.Bounds()) {
						fmt.Printf("Skipping '%s', file is not dog resolution\n", path)
						return nil
					}

					sizeOfMessage := steganography.GetMessageSizeFromImage(img)

					fmt.Println("sizeOfMessage ====", sizeOfMessage)
					fmt.Println("info Size ====", info.Size())

					hiddenFile := steganography.Decode(sizeOfMessage, img)

					deFile, err := gcm.Decrypt([]byte("RgUkXp2r5u8x/A?D(G+KbPeShVmYq3t6"), hiddenFile, 12)
					if err != nil {
						panic(err)
					}

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
