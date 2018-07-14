FROM sdqali.in/go/todo/base:31f7924

WORKDIR /app

ADD out/docker/todo-server /app/
ADD dbdo.sh /app/
ADD migrations /app/migrations

CMD ["sh", "-c", "./dbdo.sh && ./todo-server --store=pg"]
