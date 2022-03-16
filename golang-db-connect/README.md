# golang web app connect with postgres sql

## install postgres and pgadmin on mac
Follow this instruction here
https://www.enterprisedb.com/postgres-tutorials/installation-postgresql-mac-os

## create a table with pgadmin
use the pgadmin tool to create a test table and some columns and some test data

## connect db with golang webapp
- use the golang driver package recommended [here](https://github.com/golang/go/wiki/SQLDrivers)

- I am using `https://github.com/jackc/pgx`

- run `go get github.com/jackc/pgx/v4`
`

## how to write sql queries with golang
https://pkg.go.dev/database/sql

