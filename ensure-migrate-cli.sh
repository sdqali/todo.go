#!/bin/sh

mkdir $OUT_DIR
go get -u -d github.com/golang-migrate/migrate/cli
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
cd $GOPATH/src/github.com/golang-migrate/migrate/cli
dep ensure
cd -
go build -tags 'postgres' -o $OUT_DIR/migrate github.com/golang-migrate/migrate/cli

