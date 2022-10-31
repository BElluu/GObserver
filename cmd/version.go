package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check GObserver version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GObserver v1.0 / 01-11-2022")
	},
}
