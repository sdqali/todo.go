FROM iron/go

WORKDIR /app

ADD out/docker/todo-server /app/
ADD out/docker/migrate /app/
ADD migrations /app/migrations

CMD ["sh", "-c", "./migrate && ./todo-server --store=db"]
