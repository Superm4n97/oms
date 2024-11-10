# auth server
### envs
```shell
# postgres db connection string, default: "postgres://postgres:postgres@localhost:5432/oms"
DB_CONN=""
REDIS_URL=""
REDIS_USERNAME=""
REDIS_PASSWORD=""
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
go run ./oms/main.go
```