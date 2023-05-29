# [PUT] Update an existing User  

To update an existing user according to body parameters

## URL

| ** Method **    | PUT                          | 
| --------------- | ---------------------------- | 
| ** Structure ** | `/api/v1/user/update/{id}`   |


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
| first_name | string  | false    |                 | a user's first name                                                                   |
| last_name  | string  | false    |                 | a user's last name                                                                    |
| email      | string  | false    |                 | a user's email using for login to application                                         |
| password   | string  | false    |                 | a user's password using for login to application                                      |
| is_active  | boolean | false    |                 | In order to activate this existing user, set to be `true` to activate. Otherwise, `false`  |

## Sample Request(s) 
### A sample request to update a password where user ID = 1 
```
url = /api/v1/user/update/1
```
```json
{
  "password" : "1234"
}
```

### A sample request to update user information where user ID = 1 
```
url = /api/v1/user/update/1
```
```json
{
  "first_name" : "Velma",
  "last_name" : "Barry"
}
```

## Sample Response(s)
### A success response updating a password
```json
{
  "status" : "SUCCESS",
  "code" : 200
}
```

### A success response updating user information
```json
{
  "status" : "SUCCESS",
  "code" : 200,
  "data" : {
    "first_name" : "Velma",
    "last_name" : "Barry"
  }
}
```

### An error response (case: no parameter in a request body)
```json
{
  "status" : "ERROR",
  "code" : 400,
  "message" : "Bad Request"
}
```
