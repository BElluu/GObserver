package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(watchCmd)
}

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Show table with machines",
	Long:  "Show table with machines. If you use flags, you can filter and sort this table",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tu będzie tabelka")
	},
}
