package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gobserver/data"
	"gobserver/utils"
	"io/ioutil"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add server to your collection",
	Long:  "Add server to your collection. You can use tags to group your servers.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding server to your collection")
		addServer(args[0], args[1], args[2:]...)
	},
}

func addServer(name, ipAddress string, tags ...string) {
	var tagSlice []string
	var dataServers = data.MyServers.Server

	for _, tag := range tags {
		tagSlice = append(tagSlice, tag)
	}

	slug := utils.UniqueSlug()
	if slug == "" {
		log.Fatal("Something went wrong...")
	}

	onlineStatus := utils.PingTarget(ipAddress)

	newServer := data.ServerDetails{
		Id:             slug,
		Name:           name,
		IpAddress:      ipAddress,
		Online:         onlineStatus,
		LastTimeOnline: time.Now().Format("02-01-2006 15:01:05"),
		Tags:           tagSlice,
	}

	file, _ := ioutil.ReadFile("/home/bartek/Programming/GObserver/data/servers.json")
	if len(file) != 0 {
		err := json.Unmarshal(file, &dataServers)
		if err != nil {
			return
		}

		var xs []map[string]interface{}
		err = json.Unmarshal(file, &xs)
		if err != nil {
			return
		}
		for _, field := range xs {
			if field["Name"] == name {
				log.Fatalf("Server with %s exists", name)
			}
			if field["IpAddress"] == ipAddress {
				log.Fatalf("Server with %s exists", ipAddress)
			}
		}
	}

	data.MyServers.Server = append(dataServers, newServer)
	dataBytes, _ := json.MarshalIndent(data.MyServers.Server, "", " ")
	_ = ioutil.WriteFile("/home/bartek/Programming/GObserver/data/servers.json", dataBytes, 0644)
}
