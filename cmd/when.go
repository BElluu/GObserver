package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(whenCmd)
}

var whenCmd = &cobra.Command{
	Use:   "when",
	Short: "Get information when last time server was online",
	Long:  "Get information when last time server was online",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Host: ABC. Was online 22.01.2022 13:51:11")
	},
}
