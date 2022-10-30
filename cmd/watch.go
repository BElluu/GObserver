package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gobserver/data"
	"gobserver/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

const ESC = 27

func init() {
	rootCmd.AddCommand(watchCmd)
}

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Show table with machines",
	Long:  "Show table with machines. If you use flags, you can filter and sort this table",
	Run: func(cmd *cobra.Command, args []string) {
		err := showServers()
		if err != nil {
			return
		}
	},
}

func showServers() error {
	std := os.Stdout
	format := "%s\t%s\t%s\t%s\t%s\t%s\n"
	clear := fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)
	_, _ = fmt.Fprint(std, strings.Repeat(clear, 1000))
	writer := tabwriter.NewWriter(std, 10, 1, 5, ' ', 0)

	lastModifyTable := time.Now().Format("2006-01-02 15:04:05")
	firstLoad := true
	ticker := time.NewTicker(15 * time.Second)

	for {
		if !firstLoad {
			select {
			case <-ticker.C:
				updateOnlineStatus()
			}
		}

		if !firstLoad && getLastModifyServers() <= lastModifyTable {
			continue
		}

		absPath, _ := filepath.Abs("data/servers.json")
		file, _ := ioutil.ReadFile(absPath)
		servers := data.MyServers.Server
		err := json.Unmarshal(file, &servers)
		if err != nil {
			return err
		}
		//numberOfRecords := len(servers)

		if !firstLoad {
			_, _ = fmt.Fprint(std, strings.Repeat(clear, 1000))
		}
		_, err = fmt.Fprintf(writer, format, "Id", "Name", "Ip Address", "Online", "Last online", "Tags")
		if err != nil {
			return err
		}
		for _, p := range servers {
			_, err := fmt.Fprintf(writer, format, p.Id, p.Name, p.IpAddress, strconv.FormatBool(p.Online), p.LastTimeOnline, p.Tags)
			if err != nil {
				return err
			}
		}

		err = writer.Flush()
		if err != nil {
			return err
		}

		lastModifyTable = time.Now().Format("2006-01-02 15:04:05")
		firstLoad = false

	}
}

func getLastModifyServers() string {
	absPath, _ := filepath.Abs("data/servers.json")
	file, err := os.Stat(absPath)
	if err != nil {
		fmt.Println(err)
	}
	return file.ModTime().Format("2006-01-02 15:04:05")
}

func updateOnlineStatus() {
	absPath, _ := filepath.Abs("data/servers.json")
	file, _ := ioutil.ReadFile(absPath)
	fields := make([]map[string]interface{}, 0)
	err := json.Unmarshal(file, &fields)
	if err != nil {
		log.Fatal(err)
	}
	for _, field := range fields {
		isOnline := utils.PingTarget(field["IpAddress"].(string))
		if isOnline != field["Online"] {
			field["Online"] = isOnline
		}
		if isOnline {
			field["LastTimeOnline"] = time.Now().Format("02-01-2006 15:01:05")
		}
	}
	out, _ := json.MarshalIndent(fields, "", "  ")
	_ = ioutil.WriteFile(absPath, out, 0644)
}
