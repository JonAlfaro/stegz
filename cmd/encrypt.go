package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alistanis/goenc/aes/gcm"
	"github.com/auyer/steganography"
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
		// pass := args[0]
		inFile, _ := os.Open("dog.png")   // opening file
		reader := bufio.NewReader(inFile) // buffer reader
		img, _ := png.Decode(reader)
		defer inFile.Close()
		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {

				if err != nil {
					return err
				}

				if strings.HasPrefix(filepath.Base(path), "psr") && strings.HasSuffix(filepath.Base(path), ".png") {
					fmt.Println("hidden file found ", path)
					pFile, err := ioutil.ReadFile(path)
					if err != nil {
						log.Fatal(err)
					}

					enFile, err := gcm.Encrypt([]byte("RgUkXp2r5u8x/A?D(G+KbPeShVmYq3t6"), pFile, 12)
					if err != nil {
						panic(err)
					}

					w := new(bytes.Buffer)                     // buffer that will recieve the results
					err = steganography.Encode(w, img, enFile) // Encode the message into the image
					if err != nil {
						log.Printf("Error Encoding file %v", err)
						panic(err)
					}

					outFile, _ := os.Create(path) // create file
					defer outFile.Close()
					w.WriteTo(outFile) // write buffer to it

				}

				return nil
			})

		if err != nil {
			log.Println(err)
		}
	},
}
