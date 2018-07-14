FROM sdqali.in/go/todo/base:57b0216

WORKDIR /app

ADD out/docker/todo-server /app/
ADD migrations /app/migrations

CMD ["sh", "-c", "./migrate && ./todo-server --store=db"]
