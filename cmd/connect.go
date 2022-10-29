package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gobserver/data"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to server by SSH",
	Long:  "Connect to server by SSH. Your public key must be on server.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := sshConnection(args[0], args[1], args[2])
		if err != nil {
			log.Fatal("Something went wrong...")
		}
	},
}

func sshConnection(user, server, port string) error {
	srv := &data.ServerDetails{Id: server, IpAddress: server, Name: server}
	absPath, _ := filepath.Abs("data/servers.json")
	file, _ := ioutil.ReadFile(absPath)

	var xs []map[string]interface{}
	err := json.Unmarshal(file, &xs)
	if err != nil {
		return err
	}
	for _, field := range xs {
		if field["Name"] == srv.Name || field["IpAddress"] == srv.IpAddress || field["Id"] == srv.Id {
			ip := field["IpAddress"]
			if port == "" {
				port = "22"
			}

			var host = fmt.Sprintf("%s:%s", ip, port)

			fmt.Println("Enter password: ")
			password, err := terminal.ReadPassword(0)

			sshConfig := &ssh.ClientConfig{
				User: user,
				Auth: []ssh.AuthMethod{ssh.Password(string(password))},
			}
			sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

			client, err := ssh.Dial("tcp", host, sshConfig)
			if err != nil {
				return err
			}
			input := bufio.NewReader(os.Stdin)
			fmt.Println("Connected to: " + server)
			for {
				session, err := client.NewSession()
				if err != nil {
					return err
				}
				session.Stdout = os.Stdout
				command, _ := input.ReadString('\n')
				command = strings.TrimRight(command, "\r\n")

				if command == "exit" {
					err := session.Close()
					if err != nil {
						return err
					}
					break
				} else {
					err := session.Run(command)
					if err != nil {
						return err
					}
				}
			}
			return nil
		}
	}
	return nil
}
