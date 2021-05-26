all: test build

vendor:
	go mod vendor

test:
	go test github.com/hekonsek/vrs/vrs

build: vendor
	go build -o out/vrs main/*.go