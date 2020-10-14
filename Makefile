all: test build

test:
	go test github.com/hekonsek/ver

build:
	go build -o out/versioon main/*.go

docker-build: build
	docker build out -t hekonsek/versioon

docker-push: docker-build
	docker push hekonsek/versioon