#!/bin/sh

./migrate -database $DATABASE_URL -source file://migrations up
