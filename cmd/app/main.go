package main

import (
	"log"
	"strconv"
	"strings"

	aw "github.com/deanishe/awgo"
	"go.deanishe.net/env"
)

var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

func run() {
	args := wf.Args()

	log.Printf("alfred_workflow_bundleid = '%v'\n", env.Get("alfred_workflow_bundleid"))
	log.Printf("alfred_workflow_cache = '%v'\n", env.Get("alfred_workflow_cache"))
	log.Printf("alfred_workflow_data = '%v'\n", env.Get("alfred_workflow_data"))

	log.Printf("Arguments count = %v\n", len(args))

	query := args[0]
	values := strings.Fields(query)

	results := []string{}
	for i := 0; i < len(values); i++ {
		n, err := strconv.ParseInt(values[i], 16, 64)
		if err != nil {
			results = append(results, "ERR")
		} else {
			results = append(results, strconv.FormatInt(n, 10))
		}
	}

	wf.NewItem(strings.Join(results, " "))

	wf.SendFeedback()

}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
