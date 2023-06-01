# [GET] Get User Information 

To get user information such as first name, last name, email, etc.

## URL

| ** Method **    | GET                       | 
| --------------- | ------------------------- | 
| ** Structure ** | `/api/v1/user`            |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                    |
| ------------------- | :--------: | :-------: | ------------------------------ |
| Authorization       | string     | true      | A bearer token is required     |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description      |
| ---------- | ------- | -------- | --------------- | ----------------- |
|            |         |          |                 |                   |


## Sample Request(s) 

### A sample request according to user ID from authorization token
```
url = /api/v1/user
```

## Sample Response(s) 
### A success response
HTTP status 200 OK
```json
{
  "data": {
    "id" : "1",
    "first_name" : "Gloria",
    "last_name" : "Bonner",
    "email" : "gloria.bonner@gmail.com",
    "is_active" : "true"
  }
}
```

### A success response (case: data doesn't exist)
HTTP status 404 Not Found
```json
{
  "error": "user 1 does not exist"
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
