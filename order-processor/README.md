# auth server
### envs
```shell
REDIS_PASSWORD="redis"
REDIS_URL="localhost:6379"
REDIS_USERNAME="default"
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