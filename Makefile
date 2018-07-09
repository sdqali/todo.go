.PHONY : test build

build:
	go build github.com/sdqali/todo/cmd/todo

test:
	go test github.com/sdqali/todo/test
