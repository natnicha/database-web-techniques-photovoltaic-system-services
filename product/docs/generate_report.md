# [POST] Generate a Photovoltaic System Report

To generate a photovoltaic report according to the product for a specific user in authorization token. This API only generate a report and place on a system. It will not send out an email for a report. This API is designed to work together with `/api/v1/project/generate-report/{id}`

## URL

| ** Method **    | POST                                   | 
| --------------- | -------------------------------------- | 
| ** Structure ** | `/api/v1/product/generate-report/{id}` |


## Path Parameters

| Key       | Type      | Required     | Description                         |
| --------- | :-------: | :----------: | ----------------------------------- |
| id        | integer   | true         | a product ID to generate a report   |


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
url = /api/v1/product/generate-report/1
```

## Sample Response(s)
### A success Response
HTTP status 202 Accepted
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
