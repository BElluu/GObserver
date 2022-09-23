package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gobserver/data"
	"gobserver/utils"
	"io/ioutil"
	"time"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add server to your collection",
	Long:  "Add server to your collection. You can use tags to group your servers.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tu będzie dodawanie serwera")
		addServer("Test", "192.123.111.12", "Quadro", "Vivo")
	},
}

func addServer(name, ipAddress string, tags ...string) {

	var tagSlice []string

	for _, tag := range tags {
		tagSlice = append(tagSlice, tag)
	}
	slug := utils.UniqueSlug()
	if slug == "" {
		panic("Something went wrong...")
	}

	newServer := data.ServerDetails{
		Id:             slug,
		Name:           name,
		IpAddress:      ipAddress,
		Online:         false,
		LastTimeOnline: time.Now().Format("02-01-2006 15:01:05"),
		Tags:           tagSlice,
	}

	file, _ := ioutil.ReadFile("/home/bartek/Programming/GObserver/data/servers.json")
	err := json.Unmarshal(file, &data.MyServers.Servers)
	if err != nil {
		return
	}
	data.MyServers.Servers = append(data.MyServers.Servers, newServer)

	dataBytes, _ := json.MarshalIndent(data.MyServers.Servers, "", " ")

	_ = ioutil.WriteFile("/home/bartek/Programming/GObserver/data/servers.json", dataBytes, 0644)
}
