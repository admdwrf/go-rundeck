package main

import (
	"fmt"
	"os"

	"github.com/admdwrf/go-rundeck/sdk/v18"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	jobid  = kingpin.Arg("jobid", "The id of the job to export").Required().String()
	format = kingpin.Flag("format", "Format to export").Default("xml").Enum("xml", "yaml")
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	res, err := client.ExportJob(*jobid, *format)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf(res)
		os.Exit(0)
	}
}
