package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export your servers collection to JSON file",
	Long:  "Export your servers collection to JSON file. Exported file you can import on other machine using gobs import command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exporting collection")
	},
}
