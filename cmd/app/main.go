package main

import (
	"fmt"
	"net/url"
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

	unit := args[0]
	unit = strings.ToLower(unit)

	query := args[1]

	switch unit {
	case "bin":
		convertBinUnit(query)
	case "hex":
		convertHexUnit(query)
	case "url":
		convertUrl(query)

	default:
		panic(fmt.Sprintf("%v is an unknown unit", unit))
	}

	wf.SendFeedback()
}

func addItem(title string, subTitle string, copyText string) {
	item := wf.NewItem(title)
	item.Subtitle(subTitle)
	item.Arg(copyText)
	item.Valid(true)
}

func convertUrl(query string) {
	url, err := url.Parse(query)
	if err != nil {
		panic(err)
	}

	if len(url.Scheme) > 0 {
		addItem("Scheme", url.Scheme, url.Scheme)
	}

	if len(url.Host) > 0 {
		addItem("Host", url.Host, url.Host)
	}

	urlPort := url.Port()
	if len(urlPort) > 0 {
		addItem("Port", urlPort, urlPort)
	}

	if len(url.Path) > 0 {
		addItem("Path", url.Path, url.Path)
	}

	if len(url.Fragment) > 0 {
		addItem("Fragment", url.Fragment, url.Fragment)
	}

	urlQuery := url.Query()
	for k, v := range urlQuery {
		queryItem := wf.NewItem(fmt.Sprintf("Query param '%s'", k))
		queryItem.Subtitle(v[0])
		queryItem.Arg(fmt.Sprintf("%s=%s", k, v[0]))
		queryItem.Valid(true)
	}
}

func convertBinUnit(query string) {
	convertedQuery := strings.ReplaceAll(query, " ", "")

	result, err := strconv.ParseInt(convertedQuery, 2, 64)
	if err != nil {
		panic(err)
	} else {
		convertedResult := strconv.FormatInt(result, 10)

		item := wf.NewItem(convertedResult)
		item.Valid(true)
		item.Arg(convertedResult)
	}
}

func convertHexUnit(query string) {
	if query[0] == '#' {
		query = query[1:]
	}

	fields := strings.Fields(query)

	results := []string{}
	for i := 0; i < len(fields); i++ {
		f := fields[i]
		h, err := strconv.ParseInt(f, 16, 64)
		if err != nil {
			panic(err)
		} else {
			results = append(results, strconv.FormatInt(h, 10))
		}
	}

	convertedResult := strings.Join(results, " ")

	item := wf.NewItem(convertedResult)
	item.Valid(true)
	item.Arg(convertedResult)
}
