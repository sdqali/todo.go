.PHONY : test build

PROJECT="github.com/sdqali/todo"

build:
	go build -v -o out/$(PREFIX)/todo $(PROJECT)/cmd/todo
	go build -v -o out/$(PREFIX)/todo-server $(PROJECT)/server/todo-server

test:
	go test -v $(PROJECT)
	go test -v $(PROJECT)/store/json

docker:
	docker build -t sdqali.in/go/todo:latest .
	docker image prune -f

deploy:
	heroku container:push web && heroku container:release web
	docker image prune -f
