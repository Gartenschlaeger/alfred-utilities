#!/usr/bin/env zsh

export alfred_workflow_bundleid="de.kaisnet.de"
export alfred_workflow_cache="$(PWD)/cache"
export alfred_workflow_data="$(PWD)/data"

# Tests

go run cmd/app/* "bin" "010010111"

go run cmd/app/* "hex" "aa"
go run cmd/app/* "hex" "ff ee ae"
