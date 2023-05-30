# [POST] Create a New User  

To create a new user according to body parameters

## URL

| ** Method **    | POST                    | 
| --------------- | ----------------------- | 
| ** Structure ** | `/auth/register`        |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                                   |
| ------------------- | :--------: | :-------: | --------------------------------------------- |
| Content-Type        | string     | true      | Content-Type has to be `application/json`     |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description                                                                          |
| ---------- | ------- | -------- | --------------- | ------------------------------------------------------------------------------------- |
| first_name | string  | true     |                 | a user's first name                                                                   |
| last_name  | string  | true     |                 | a user's last name                                                                    |
| email      | string  | true     |                 | a user's email using for login to application                                         |
| password   | string  | true     |                 | a user's password using for login to application                                      |
| is_active  | boolean | false    | true            | In order to activate this new user, set to be `true` to activate. Otherwise, `false`  |

## Sample Request(s) 
```
url = /auth/register
```
```json
{
  "first_name" : "Gloria",
  "last_name" : "Bonner",
  "email" : "gloria.bonner@gmail.com",
  "password" : "1234",
  "is_active" : "true"
}
```

## Sample Response(s)
### A success Response
HTTP status 201 Created
```json
{
  "data": {
    "id" : "1",
    "first_name" : "Gloria",
    "last_name" : "Bonner",
    "email" : "gloria.bonner@gmail.com",
    "password" : "$2a$10$RqZ3UIVfsSM/jP/dO3.5.u2OxuJBU29YvPlYQdPg1cnTax4D8Ny7C",
    "is_active" : "true"
  }
}
```

### An error response (case: missing first name)
HTTP status 400 Bad Request
```json
null
```

### An error response (case: invalid email format)
HTTP status 400 Bad Request
```json
{
  "error": "The email is invalid format"
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```

### An error response (case: conflict email in the system)
HTTP status 409 Conflict
```json
{
  "error": "The email is already assigned in the system"
}
```
