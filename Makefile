.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v $(PROJECT)/cmd/todo

build-server:
	go build -v $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)

docker-build:
	docker run --rm -v `pwd`:/go/src/$(PROJECT) -w /go/src/$(PROJECT) iron/go:dev make build build-server
