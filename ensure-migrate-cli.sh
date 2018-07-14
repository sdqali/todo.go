#!/usr/bin/env bash

OUT_DIR="${OUT_DIR:-out}"
mkdir -p $OUT_DIR

if [ -f $OUT_DIR/migrate ]; then
  exit 0
fi

if [ ! -f $OUT_DIR/migrate.tar.gz ]; then
  wget "https://github.com/golang-migrate/migrate/releases/download/v3.3.0/migrate.linux-amd64.tar.gz" -O $OUT_DIR/migrate.tar.gz
fi

pushd $OUT_DIR
tar xf migrate.tar.gz
mv migrate.linux-amd64 migrate
popd
