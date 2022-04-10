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

	operation := args[0]
	operation = strings.ToLower(operation)

	query := args[1]

	switch operation {

	case "bin2dec":
		convertBin2Dec(query)
	case "dec2bin":
		convertDec2Bin(query)

	case "hex2dec":
		convertHex2Dec(query)
	case "dec2hex":
		convertDec2Hex(query)

	case "url":
		parseUrl(query)

	default:
		panic(fmt.Sprintf("'%s' is an unknown operation", operation))
	}

	wf.SendFeedback()
}

func addWorkflowItem(title string, subTitle string, copyText string) {
	item := wf.NewItem(title)
	item.Subtitle(subTitle)
	item.Arg(copyText)
	item.Valid(true)
}

func parseUrl(query string) {
	url, err := url.Parse(query)
	if err != nil {
		panic(err)
	}

	if len(url.Scheme) > 0 {
		addWorkflowItem("Scheme", url.Scheme, url.Scheme)
	}

	if len(url.Host) > 0 {
		addWorkflowItem("Host", url.Host, url.Host)
	}

	urlPort := url.Port()
	if len(urlPort) > 0 {
		addWorkflowItem("Port", urlPort, urlPort)
	}

	if len(url.Path) > 0 {
		addWorkflowItem("Path", url.Path, url.Path)
	}

	if len(url.Fragment) > 0 {
		addWorkflowItem("Fragment", url.Fragment, url.Fragment)
	}

	urlQuery := url.Query()
	for k, v := range urlQuery {
		queryItem := wf.NewItem(fmt.Sprintf("Query param '%s'", k))
		queryItem.Subtitle(v[0])
		queryItem.Arg(fmt.Sprintf("%s=%s", k, v[0]))
		queryItem.Valid(true)
	}
}

func convertBin2Dec(query string) {
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

func convertDec2Bin(query string) {
	n, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		panic(err)
	}

	b := strconv.FormatInt(n, 2)

	// pad string with zeros to have 8 chacters per byte
	d := len(b) % 8
	if d != 0 {
		pl := strings.Repeat("0", 8-d)
		b = pl + b
	}

	item := wf.NewItem(b)
	item.Valid(true)
	item.Arg(b)
}

func convertHex2Dec(query string) {
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

func convertDec2Hex(query string) {
	n, err := strconv.Atoi(query)
	if err != nil {
		panic(err)
	}

	h := fmt.Sprintf("%x", n)

	item := wf.NewItem(h)
	item.Valid(true)
	item.Arg(h)
}
