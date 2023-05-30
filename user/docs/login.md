# [POST] Login a User  

To login a user account to a system and will be responded an JWT session token  

## URL

| ** Method **    | POST                    | 
| --------------- | ----------------------- | 
| ** Structure ** | `/auth/login`           |


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
| email      | string  | true     |                 | a user's email using for login to application                                         |
| password   | string  | true     |                 | a user's password using for login to application                                      |

## Sample Request(s) 
```
url = /auth/login
```
```json
{
  "email" : "gloria.bonner@gmail.com",
  "password" : "1234",
}
```

## Sample Response(s)
### A success Response
HTTP status 201 Created
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODU0NTE4ODEsImlhdCI6MTY4NTQ1MTg4MSwiaWQiOjF9.r6aK1T9ZpBxbkbBahDIg7qIOHLoFV3sQgvwWssqGNqE"
}
```

### An error response (case: missing email or password)
HTTP status 400 Bad Request
```json
null
```

### An error response (case: incorrect email or password)
HTTP status 401 Unauthorized
```json
{
  "error": "incorrect email or password"
}
```
