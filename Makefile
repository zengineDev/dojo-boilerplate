# Include variables from the .envrc file
include .envrc

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${DOJO_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	migrate create -seq -ext=.sql -dir database/migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	migrate -path ./database/migrations -database ${DOJO_DB_DSN} up