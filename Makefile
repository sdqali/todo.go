.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v -o out/$(PREFIX)/migrate $(PROJECT)/db/migrate
	go build -v -o out/$(PREFIX)/todo $(PROJECT)/cmd/todo
	go build -v -o out/$(PREFIX)/todo-server $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)

docker-build:
	godep save ./...
	docker run --rm -v `pwd`:/go/src/$(PROJECT) -w /go/src/$(PROJECT) iron/go:dev make PREFIX=docker build

push: docker-build
	heroku container:push web && heroku container:release web
