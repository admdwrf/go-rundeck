package main

import (
	"fmt"
	"os"

	"github.com/admdwrf/go-rundeck/sdk/v18"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	project = kingpin.Arg("project", "Project to delete").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()

	err := client.DeleteProject(*project)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Project %s deleted\n", *project)
		os.Exit(0)
	}
}
