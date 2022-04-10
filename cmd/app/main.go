package main

import (
	"fmt"
	"strconv"
	"strings"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}

func run() {
	args := wf.Args()
	if len(args) < 2 {
		panic("At least two arguments are expected")
	}

	//log.Printf("alfred_workflow_bundleid = '%v'\n", env.Get("alfred_workflow_bundleid"))
	//log.Printf("alfred_workflow_cache = '%v'\n", env.Get("alfred_workflow_cache"))
	//log.Printf("alfred_workflow_data = '%v'\n", env.Get("alfred_workflow_data"))

	//log.Printf("Arguments count = %v\n", len(args))

	unit := args[0]
	unit = strings.ToLower(unit)

	query := args[1]

	switch unit {
	case "bin":
		convertBinUnit(query)
	case "hex":
		convertHexUnit(query)

	default:
		panic(fmt.Sprintf("%v is an unknown unit", unit))
	}

	wf.SendFeedback()
}

func convertBinUnit(query string) {
	convertedQuery := strings.ReplaceAll(query, " ", "")

	result, err := strconv.ParseInt(convertedQuery, 2, 64)
	if err != nil {
		panic(err)
	} else {
		wf.NewItem(strconv.FormatInt(result, 10))
	}
}

func convertHexUnit(query string) {
	fields := strings.Fields(query)

	results := []string{}
	for i := 0; i < len(fields); i++ {
		n, err := strconv.ParseInt(fields[i], 16, 64)
		if err != nil {
			results = append(results, "ERROR")
		} else {
			results = append(results, strconv.FormatInt(n, 10))
		}
	}

	wf.NewItem(strings.Join(results, " "))
}
