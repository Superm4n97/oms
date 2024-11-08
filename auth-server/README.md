# auth server
### envs
```shell
# postgres db connection string, default: "postgres://postgres:postgres@localhost:5432/postgres"
DB_CONN=""
```
### run
```shell
# go to pgadmin and create a table with username
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(100)
    password VARCHAR(100)
);

# export envs
DB_CONN="postgres://postgres:postgres@localhost:5432/postgres"
# run server
go run ./auth-server/main.go
```