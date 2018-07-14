.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v -o out/$(PREFIX)/todo $(PROJECT)/cmd/todo
	go build -v -o out/$(PREFIX)/todo-server $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)

build-migrate-cli:
	OUT_DIR=out/$(PREFIX) ./build-migrate-cli.sh


docker-build:
	godep save ./...
	docker run --rm -v `pwd`:/go/src/$(PROJECT) -w /go/src/$(PROJECT) iron/go:dev make PREFIX=docker build build-migrate-cli

deploy: docker-build
	heroku container:push web && heroku container:release web
