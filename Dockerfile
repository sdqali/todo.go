FROM iron/go

WORKDIR /app

ADD todo-server /app/

CMD ["./todo-server"]
