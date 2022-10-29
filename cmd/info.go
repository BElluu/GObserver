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
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get informations about specfic server",
	Long:  "Get informations about specific server",
	Run: func(cmd *cobra.Command, args []string) {
		getInfo(args[0])
	},
}

func getInfo(server string) {
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
			fmt.Printf("Server ID: %s\n Server name: %s\n Server IP: %s\n Server was online last time at %s\n Online: %t\n Tags: %s\n",
				field["Id"], field["Name"], field["IpAddress"], field["LastTimeOnline"], field["Online"], field["Tags"])
			return
		}
	}
	fmt.Printf("Server %s not exists\n", server)
}
