all: test build

test:
	go test github.com/hekonsek/vrs/vrs

build:
	go build -o out/vrs main/*.go

docker-build: build
	docker build out -t hekonsek/versioon

docker-push: docker-build
	docker push hekonsek/versioon