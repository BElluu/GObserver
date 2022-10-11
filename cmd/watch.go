package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gobserver/data"
	"io/ioutil"
	"log"
	"os"
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
			log.Fatal("Something went wrong...")
		}
	},
}

func showServers() error {
	std := os.Stdout
	format := "%s\t%s\t%s\t%s\t%s\t%s\n"
	clear := fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)
	writer := tabwriter.NewWriter(std, 10, 1, 5, ' ', 0)

	lastModifyTable := time.Now().Format("2006-01-02 15:04:05")
	firstLoad := true

	_, err := fmt.Fprintf(writer, format, "Id", "Name", "Ip Address", "Online", "Last online", "Tags")
	if err != nil {
		return err
	}

	for {
		if !firstLoad && getLastModifyServers() <= lastModifyTable {
			continue
		}
		file, _ := ioutil.ReadFile("/home/bartek/Programming/GObserver/data/servers.json")
		servers := data.MyServers.Server
		err := json.Unmarshal(file, &servers)
		if err != nil {
			return err
		}
		numberOfRecords := len(servers)

		if !firstLoad {
			_, _ = fmt.Fprint(std, strings.Repeat(clear, numberOfRecords))
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
	file, err := os.Stat("/home/bartek/Programming/GObserver/data/servers.json")
	if err != nil {
		fmt.Println(err)
	}
	return file.ModTime().Format("2006-01-02 15:04:05")
}
