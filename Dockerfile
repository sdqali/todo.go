FROM iron/go:dev AS builder
LABEL stage=todo-intermediate
COPY . /go/src/github.com/sdqali/todo
WORKDIR /go/src/github.com/sdqali/todo
RUN make PREFIX=docker build

FROM migrate/migrate:v3.3.0

ENTRYPOINT []
WORKDIR /app

COPY --from=builder /go/src/github.com/sdqali/todo/out/docker/todo-server /app/
ADD dbdo.sh /app/
ADD migrations/postgres /app/migrations/postgres

CMD ["sh", "-c", "./dbdo.sh && ./todo-server --store=pg"]
