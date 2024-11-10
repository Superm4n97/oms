# oms
## run
### install redis
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install redis -n oms bitnami/redis --set auth.enabled=false
```
### install oms chart
* Clone this GitHub repository
* run the following helm command to oms system
    ```shell
    helm upgrade -i oms ./installer -n oms
    ```
### create postgres tables
```shell
#exec in postgres database pod
kubectl exec -it -n oms oms-postgres-586957879f-b92vt bash

#open psql console
psql -U postgres -d oms

#user table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(100)
);
#order table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    description VARCHAR(100)
);
```

### install kong
```shell
helm repo add kong https://charts.konghq.com
helm repo update
helm install kong kong/ingress -n oms --create-namespace --set gateway.image.repository=kong --set gateway.image.tag="3.8.0"

#port forward gateway proxy to use it locally 
kubectl port-forward -n oms service/kong-gateway-proxy 8000:80
```

### uninstall
* uninstall redis
  ```shell
  helm uninstall -n oms redis
  ```
* uninstall kong proxy server
  ```shell
  helm uninstall -n oms kong
  ```
* uninstall oms chart
  ```shell
  helm uninstall oms -n oms
  ```
## design
![Creation](./assets/img/design.jpg)

## Auth Server
Auth Server will provide you user authentication related information, such as: 
* create user
* list users
* get token

### create new user
`/auth/user [POST]` api endpoint will create a new user. Payload of the parameter will be

```json
{
  "username": "foobar",
  "password": "1234"
}
```
### get users
`/auth/user [GET]` api endpoint will get the list of users.

### get token
`/auth/user/:<username> [GET]` api endpoint will generate the token for user _**username**_.

## Order Management Server

### create new user
`/product/order [POST]` api endpoint will create a new user. Payload of the parameter will be

```json
{
  "username": "foobar",
  "description": "order for food"
}
```