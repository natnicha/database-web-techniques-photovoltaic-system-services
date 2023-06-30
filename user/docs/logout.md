# [POST] Logout a User  

To logout a user account from a system and will be responded an JWT session token  

## URL

| ** Method **    | POST                    | 
| --------------- | ----------------------- | 
| ** Structure ** | `/api/v1/user/logout`   |


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
| Authorization       | string     | true      | A bearer token is required                    |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description                              |
| ---------- | ------- | -------- | --------------- | ----------------------------------------- |
|            |         |          |                 |                                           |

## Sample Request(s) 
```
url = /api/v1/user/logout
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODU0NTE4ODEsImlhdCI6MTY4NTQ1MTg4MSwiaWQiOjF9.r6aK1T9ZpBxbkbBahDIg7qIOHLoFV3sQgvwWssqGNqE"
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
