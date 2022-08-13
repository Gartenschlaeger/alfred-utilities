test: export alfred_workflow_bundleid = de.kaisnet.de
test: export alfred_workflow_cache    = $(PWD)/cache
test: export alfred_workflow_data     = $(PWD)/data

clean:
	rm -rf cache
	rm -rf data
	rm -rf dist

build: clean
	go build -o dist/conv cmd/app/*

test:
	@@echo $$alfred_workflow_cache
	go run cmd/app/main.go $$operation $$query
