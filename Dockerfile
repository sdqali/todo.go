FROM iron/go

WORKDIR /app

ADD out/todo-server /app/
ADD out/migrate /app/
ADD migrations /app/migrations

CMD ["sh", "-c", "./migrate && ./todo-server --store=db"]
