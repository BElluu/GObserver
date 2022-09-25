package cmd

import (
	"encoding/json"
	"fmt"
	"gobserver/data"
	"io/ioutil"
	"log"

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
		fmt.Println("Usuwanie servera...")
		deleteServer(args[0])
	},
}

func deleteServer(idServer string) {
	elementToDelete := &data.ServerDetails{Id: idServer}

	_, _ = json.Marshal(elementToDelete)
	file, _ := ioutil.ReadFile("/home/bartek/Programming/GObserver/data/servers.json")
	fields := make([]map[string]interface{}, 0)
	err := json.Unmarshal(file, &fields)
	if err != nil {
		log.Fatal(err)
	}

	length := len(fields)
	for index, field := range fields {
		if field["Id"] == elementToDelete.Id {
			if index == length-1 {
				fields = fields[0:index]
			} else {
				fields = append(fields[0:index], fields[index+1:]...)
			}
		}
	}

	out, _ := json.MarshalIndent(fields, "", "  ")
	_ = ioutil.WriteFile("/home/bartek/Programming/GObserver/data/servers.json", out, 0644)
}
