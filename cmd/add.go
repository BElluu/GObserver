package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"gobserver/data"

	"github.com/spf13/cobra"
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
		addServer()
	},
}

func addServer() {

	tg := []string{}
	tg = append(tg, "asb", "aaaa")

	newServer := data.ServerDetails{
		Id:             4,
		Name:           "Testowya12",
		IpAddress:      "192.168.0.2",
		Online:         true,
		LastTimeOnline: time.Now(),
		Tags:           tg,
	}
	//data.MyServers.Servers.Server = newServer
	data.MyServers.Servers = append(data.MyServers.Servers, newServer)
	data.MyServers.Servers = append(data.MyServers.Servers, newServer)
	// fmt.Println(len(data.MyServers.Items))
	file, _ := json.MarshalIndent(data.MyServers, "", " ")
	_ = ioutil.WriteFile("C:\\Programmer\\Gobs\\data\\servers.json", file, 0644)
	//return &data.MyServers
	// srv := data.Servers{}

	// srv.Items = append(srv.Items, newServer)

	// fmt.Println(srv.SrvItems())
	// return &srv

	// newServer := &data.ServerItem{
	// 	Id:             1,
	// 	Name:           "Testowy",
	// 	IpAddress:      "192.168.0.1",
	// 	Online:         true,
	// 	LastTimeOnline: time.Now(),
	// }
	// data.Srv.Items = append(data.Srv.Items, *newServer)
	//fmt.Println(newServer)
	//fmt.Println(data.Srv.Items)
}
