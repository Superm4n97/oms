# oms
## design
![Creation](./assets/img/design.jpg)

## Auth Server
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