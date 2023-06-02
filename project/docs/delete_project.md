# [DELETE] Delete a Project  

To delete a project for a specific user in authorization token

## URL

| ** Method **    | DELETE                        | 
| --------------- | ----------------------------- | 
| ** Structure ** | `/api/v1/project/delete/{id}` |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
| id        | integer   | true         | a project ID to delete          |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                                 |
| ------------------- | :--------: | :-------: | ------------------------------------------- |
| Content-Type        | string     | true      | Content-Type has to be `application/json`   |
| Authorization       | string     | true      | A bearer token is required                  |


## Body Parameters

| Field Name   | Type     | Required | Default Value   |  Description                  |
| ------------ | -------- | -------- | --------------- | ----------------------------- |
|              |          |          |                 |                               |


## Sample Request(s) 
```
url = /api/v1/project/delete/1
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
null
```

### An error response (case: No specified product ID or a project ID doesn't belong to user ID in authorization token)
HTTP status 409 Conflict
```json
{
  "error": "No specified product ID or a project ID doesn't belong to a user ID"
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
