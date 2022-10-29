package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gobserver/data"
	"io/ioutil"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(whenCmd)
}

var whenCmd = &cobra.Command{
	Use:   "when",
	Short: "Get information when last time server was online",
	Long:  "Get information when last time server was online",
	Run: func(cmd *cobra.Command, args []string) {
		whenWasOnline(args[0])
	},
}

func whenWasOnline(server string) {
	checkServer := &data.ServerDetails{Id: server, IpAddress: server, Name: server}
	absPath, _ := filepath.Abs("data/servers.json")
	file, _ := ioutil.ReadFile(absPath)

	var xs []map[string]interface{}
	err := json.Unmarshal(file, &xs)
	if err != nil {
		return
	}
	for _, field := range xs {
		if field["Name"] == checkServer.Name || field["IpAddress"] == checkServer.IpAddress || field["Id"] == checkServer.Id {
			fmt.Printf("Server was online last time at %s\n", field["LastTimeOnline"])
			return
		}
	}
	fmt.Printf("Server %s not exists\n", server)
}
