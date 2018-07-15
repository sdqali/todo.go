FROM migrate/migrate:v3.3.0

ENTRYPOINT []
WORKDIR /app

ADD out/docker/todo-server /app/
ADD dbdo.sh /app/
ADD migrations/postgres /app/migrations/postgres

CMD ["sh", "-c", "./dbdo.sh && ./todo-server --store=pg"]
