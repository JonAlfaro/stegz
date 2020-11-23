package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "Stegz is a tool for recursive LSB steganography / encryption on the working directory",
		Long:  `Stegz is a tool for recursive LSB steganography / encryption on the working directory`,
	}
)

// Execute ....
func Execute() error {
	return rootCmd.Execute()
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
