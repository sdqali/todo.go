FROM iron/go

WORKDIR /app

ADD todo-server /app/
ADD migrate /app/
ADD migrations /app/migrations

CMD ["sh", "-c", "./migrate && ./todo-server"]
