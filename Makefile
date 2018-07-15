.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v -o out/$(PREFIX)/todo $(PROJECT)/cmd/todo
	go build -v -o out/$(PREFIX)/todo-server $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)

docker:
	docker build -t sdqali.in/go/todo:latest .

deploy:
	heroku container:push web && heroku container:release web
