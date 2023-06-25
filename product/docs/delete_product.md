# [DELETE] Delete a Product  

To delete a product for a specific user in authorization token

## URL

| ** Method **    | DELETE                        | 
| --------------- | ----------------------------- | 
| ** Structure ** | `/api/v1/product/delete/{id}` |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
| id        | integer   | true         | a product ID to delete          |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                                                                   | Permission         |
| ------------------- | :--------: | :-------: | ----------------------------------------------------------------------------- | ------------------ |
| Authorization       | string     | false     | A bearer token is required for external access                                | external only      |
| api-key             | string     | false     | Instead of Authorization, internal requests required only api-key and user-id | internal only      |
| user-id             | string     | false     | Instead of Authorization, internal requests required only api-key and user-id | internal only      |


## Body Parameters

| Field Name   | Type     | Required | Default Value   |  Description                  |
| ------------ | -------- | -------- | --------------- | ----------------------------- |
|              |          |          |                 |                               |


## Sample Request(s) 
```
url = /api/v1/product/delete/1
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
null
```

### An error response (case: a product ID doesn't belong to user ID in authorization token)
HTTP status 409 Conflict
```json
{
  "error": "a product ID doesn't belong to a user ID"
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
