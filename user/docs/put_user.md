# [PUT] Update an Existing User  

To update an existing user according to body parameters

## URL

| ** Method **    | PUT                          | 
| --------------- | ---------------------------- | 
| ** Structure ** | `/api/v1/user/update`        |


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
| Authorization       | string     | true      | A bearer token is required                    |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description                                                                          |
| ---------- | ------- | -------- | --------------- | ------------------------------------------------------------------------------------- |
| first_name | string  | true     |                 | a user's first name                                                                   |
| last_name  | string  | true     |                 | a user's last name                                                                    |
| email      | string  | true     |                 | a user's email using to login to an application                                         |
| password   | string  | true     |                 | a user's password using to login to an application                                      |
| is_active  | boolean | false    | true            | In order to activate this new user, set to be `true` to activate. Otherwise, `false`  |

## Sample Request(s) 
### A sample request to update a password according to user ID from authorization token
```
url = /api/v1/user/update
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
### A success response updating a password
HTTP status 200 OK
```json
{
  "id": "1",
  "first_name" : "Gloria",
  "last_name" : "Bonner",
  "email" : "gloria.bonner@gmail.com",
  "password" : "1234",
  "is_active" : "true",
}
```

### A success response updating user information
HTTP status 200 OK
```json
null
```

### An error response (case: no parameter in a request body)
HTTP status 400 Bad Request
```json
null
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
