package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

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
		export(args[0])
	},
}

func export(destination string) {
	absPath, _ := filepath.Abs("data/servers.json")
	serversJson, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatalln("Something went wrong...")
		return
	}

	err = ioutil.WriteFile(destination+".json", serversJson, 0644)
	if err != nil {
		log.Fatalf("Error creating %s", destination)
		return
	}
}
