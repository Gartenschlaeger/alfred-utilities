#!/usr/bin/env zsh

export alfred_workflow_bundleid="de.kaisnet.de"
export alfred_workflow_cache="$(PWD)/cache"
export alfred_workflow_data="$(PWD)/data"

# Tests
#go run cmd/app/* "bin2dec" "010010111"
#go run cmd/app/* "hex2dec" "aa"
#go run cmd/app/* "hex2dec" "ff ee ae"
go run cmd/app/* "hex2dec" "#ffffff"

#go run cmd/app/* "url" "https://www.test.de:3000/test/123?a=test-1&b=test+test+123+usw#test-blubb"
