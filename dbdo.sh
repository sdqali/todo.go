#!/bin/sh

./pg_migrate -database $DATABASE_URL -source file://migrations/postgres up
