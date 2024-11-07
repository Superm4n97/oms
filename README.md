# oms
## design
![Creation](./assets/img/design.jpg)

## Auth Server
### create new user
`/auth/user [POST]` api endpoint will create a new user. Payload of the parameter will be

```json
{
  "firstName": "foo",
  "lastName": "bar",
  "username": "foobar",
  "password": "1234"
}
```

## OMS Server
