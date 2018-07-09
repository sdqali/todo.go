.PHONY : test

build:
	go build todo/cmd/todo

test:
	go test todo/test
