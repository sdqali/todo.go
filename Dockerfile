FROM iron/go

WORKDIR /app

ADD todo-server /app/
ADD migrate /app/

CMD ["./migrate && ./todo-server"]
