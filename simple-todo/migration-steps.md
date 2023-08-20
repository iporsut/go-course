# Go migration

## install migrate command

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

or

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
```

## create database

```
psql -h localhost -U postgres -w -c "create database simpletodo;"
```

```
export POSTGRESQL_URL='postgres://weerasak@localhost:5432/simpletodo?sslmode=disable'
```

## create new migration

```
migrate create -ext sql -dir db/migrations -seq create_todos_table
```

## add create todos migration up script

```
CREATE TABLE IF NOT EXISTS todos(
   id serial PRIMARY KEY,
   title VARCHAR (200) NOT NULL,
   created_at timestamp without time zone,
   completed_at timestamp without time zone,
);
```

## add drop migration down script

```
DROP TABLE IF EXISTS todos;
```

## run migration

```
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
