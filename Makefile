.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v -o out/$(PREFIX)/todo $(PROJECT)/cmd/todo
	go build -v -o out/$(PREFIX)/todo-server $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)

ensure-migrate-cli:
	OUT_DIR=out/$(PREFIX) ./ensure-migrate-cli.sh

docker-deps:
	godep save ./...
	docker run --rm -v `pwd`:/go/src/$(PROJECT) -w /go/src/$(PROJECT) iron/go:dev make PREFIX=docker build

docker-base:
	docker run --rm -v `pwd`:/go/src/$(PROJECT) -w /go/src/$(PROJECT) iron/go:dev make PREFIX=docker ensure-migrate-cli
	docker build -t sdqali.in/go/todo/base:`git rev-parse --short HEAD` -f Dockerfile.base .

docker: docker-deps
	docker build -t sdqali.in/go/todo:`git rev-parse HEAD` .

deploy: docker-deps
	heroku container:push web && heroku container:release web
