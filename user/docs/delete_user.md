# [DELTE] Delete an Existing User  

To delete an existing user by authorization token

## URL

| ** Method **    | DELETE                       | 
| --------------- | ---------------------------- | 
| ** Structure ** | `/api/v1/user/delete`        |


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
```
url = /api/v1/user/delete
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
null
```

### An error response (case: user ID doesn't exist)
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
