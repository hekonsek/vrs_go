all: test build

vendor:
	go mod vendor

fmt:
	go fmt main/*.go
	go fmt vrs/*.go

test:
	go test github.com/hekonsek/vrs/vrs

build: vendor fmt
	go build -o out/vrs main/*.go

gosec:
	gosec vrs main