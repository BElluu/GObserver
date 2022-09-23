package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to server by SSH",
	Long:  "Connect to server by SSH. Your public key must be on server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connecting by SSH...")
	},
}
