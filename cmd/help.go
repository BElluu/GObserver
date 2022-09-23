package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helpCmd)
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show all available commands",
	Long:  "Show all available commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Commands list")
	},
}
