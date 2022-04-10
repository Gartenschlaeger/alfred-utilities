#!/usr/bin/env zsh

export alfred_workflow_bundleid="de.kaisnet.de"
export alfred_workflow_cache="$(PWD)/cache"
export alfred_workflow_data="$(PWD)/data"

go run cmd/app/* "hex" "fff"
