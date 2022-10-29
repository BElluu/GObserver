package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tagCmd)
}

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add or delete tags from your servers",
	Long:  "Add or delete tags from your servers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet...")
	},
}
