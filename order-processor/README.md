# auth server
### envs
```shell
REDIS_URL=""
REDIS_USERNAME=""
REDIS_PASSWORD=""
```
### run
```shell
# export envs
REDIS_URL="localhost:6379"
REDIS_USERNAME="default"
REDIS_PASSWORD="redis"
# run server
go run ./oms/main.go
```