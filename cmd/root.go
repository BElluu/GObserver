package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobs",
	Short: "GObserver is tool to monitor machines",
	Long:  "Simple and fast tool to monitor machines",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GObserver main commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
