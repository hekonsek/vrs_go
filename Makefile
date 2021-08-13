all: test build

vendor:
	go mod vendor

fmt:
	go fmt main/*.go
	go fmt vrs/*.go
	go fmt exe/*.go

test:
	go test github.com/hekonsek/vrs/vrs github.com/hekonsek/vrs/exe github.com/hekonsek/vrs/semver

build: vendor fmt
	go build -o out/vrs main/*.go

gosec:
	gosec vrs main exe semver

install: build
	sudo cp out/vrs /usr/local/bin/

release: install
	docker build -t hekonsek/vrs out
	docker push hekonsek/vrs