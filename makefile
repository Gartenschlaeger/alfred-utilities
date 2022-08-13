start:
	./test.sh

build:
	go build -o dist/conv cmd/app/*

clean:
	rm -rf cache
	rm -rf data
	rm -rf dist
