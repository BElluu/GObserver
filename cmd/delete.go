package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server from your collection",
	Long:  "Delete server from your collection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usuwanie servera...")
	},
}
