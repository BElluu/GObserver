package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gobserver/data"
	"gobserver/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

type promptContent struct {
	errorMsg        string
	label           string
	validEmptyValue bool
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add server to your collection",
	Long:  "Add server to your collection.",
	Run: func(cmd *cobra.Command, args []string) {
		addServer()
	},
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if !pc.validEmptyValue {
			if len(input) <= 0 {
				return errors.New(pc.errorMsg)
			}
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func addServer() {

	nameServerPrompt := promptContent{
		"Please provide a server name.",
		"Server name: ",
		false,
	}
	nameServer := promptGetInput(nameServerPrompt)

	ipAddressPrompt := promptContent{
		"Please provide a ip address.",
		fmt.Sprintf("Ip address for %s: ", nameServer),
		false,
	}

	ipAddress := promptGetInput(ipAddressPrompt)

	tagsPrompt := promptContent{
		"Something went wrong...",
		fmt.Sprintf("Tags(separated by a space): "),
		true,
	}

	tagsString := promptGetInput(tagsPrompt)

	tags := strings.Split(tagsString, " ")

	var dataServers = data.MyServers.Server

	slug := utils.UniqueSlug()
	if len(slug) <= 0 {
		log.Fatal("Something went wrong...")
	}

	onlineStatus := utils.PingTarget(ipAddress)

	newServer := data.ServerDetails{
		Id:             slug,
		Name:           nameServer,
		IpAddress:      ipAddress,
		Online:         onlineStatus,
		LastTimeOnline: time.Now().Format("02-01-2006 15:01:05"),
		Tags:           tags,
	}

	absPath, _ := filepath.Abs("data/servers.json")
	file, _ := ioutil.ReadFile(absPath)
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
			if field["Name"] == nameServer {
				log.Fatalf("Server with %s exists", nameServer)
			}
			if field["IpAddress"] == ipAddress {
				log.Fatalf("Server with %s exists", ipAddress)
			}
		}
	}

	data.MyServers.Server = append(dataServers, newServer)
	dataBytes, _ := json.MarshalIndent(data.MyServers.Server, "", " ")
	_ = ioutil.WriteFile(absPath, dataBytes, 0644)
}
