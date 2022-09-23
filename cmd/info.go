package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get informations about specfic server",
	Long:  "Get informations about specific server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server XYZ. IP: 127.0.0.1. etc")
	},
}
