.PHONY : test build

build:
	go build -v github.com/sdqali/todo/cmd/todo

build-server:
	go build -v github.com/sdqali/todo/server/todo-server

test:
	go test -v github.com/sdqali/todo/test
