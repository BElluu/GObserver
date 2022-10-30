package cmd

import (
	"encoding/json"
	"fmt"
	"gobserver/data"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server from your collection",
	Long:  "Delete server from your collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteServer(args[0])
		fmt.Printf("Server %s deleted if exists", args[0])
	},
}

func deleteServer(record string) {
	elementToDelete := &data.ServerDetails{Id: record, IpAddress: record, Name: record}

	_, _ = json.Marshal(elementToDelete)
	absPath, _ := filepath.Abs("data/servers.json")
	file, _ := ioutil.ReadFile(absPath)
	fields := make([]map[string]interface{}, 0)
	err := json.Unmarshal(file, &fields)
	if err != nil {
		log.Fatal(err)
	}

	length := len(fields)
	for index, field := range fields {
		if field["Id"] == elementToDelete.Id || field["IpAddress"] == elementToDelete.IpAddress || field["Name"] == elementToDelete.Name {
			if index == length-1 {
				fields = fields[0:index]
			} else {
				fields = append(fields[0:index], fields[index+1:]...)
			}
		}
	}

	out, _ := json.MarshalIndent(fields, "", "  ")
	_ = ioutil.WriteFile(absPath, out, 0644)
}
