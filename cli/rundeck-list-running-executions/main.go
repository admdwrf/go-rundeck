package main

import (
	"fmt"
	"os"
"github.com/admdwrf/go-rundeck/sdk/v18"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	project = kingpin.Arg("project", "").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	res, err := client.ListRunningExecutions(*project)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", res)
		//fmt.Printf("Job %s is %s", res.Executions[0].ID, res.Executions[0].Status)
		os.Exit(0)
	}
}
