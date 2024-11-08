# auth server
### envs
```shell
# postgres db connection string, default: "postgres://postgres:postgres@localhost:5432/oms"
DB_CONN=""
```
### run
```shell
# go to pgadmin and create a table with username
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    description VARCHAR(100)
);

# export envs
DB_CONN="postgres://postgres:postgres@localhost:5432/oms"
# run server
go run ./auth-server/main.go
```