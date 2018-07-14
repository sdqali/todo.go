.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v -o out/$(PREFIX)/todo $(PROJECT)/cmd/todo
	go build -v -o out/$(PREFIX)/todo-server $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)

ensure-migrate-cli:
	OUT_DIR=out/docker ./ensure-migrate-cli.sh

docker-deps: ensure-migrate-cli
	godep save ./...
	docker run --rm -v `pwd`:/go/src/$(PROJECT) -w /go/src/$(PROJECT) iron/go:dev make PREFIX=docker build

docker:
	docker build -t todo/`git rev-parse HEAD` .

deploy: docker-deps
	heroku container:push web && heroku container:release web
