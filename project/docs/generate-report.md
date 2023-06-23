# [POST] Generate a Photovoltaic System Report within a Project

To generate photovoltaic reports by calling to /product/generate-report/{id} according to the project ID for a specific user in authorization token and send out an email once finished generating all reports within the project

## URL

| ** Method **    | POST                                   | 
| --------------- | -------------------------------------- | 
| ** Structure ** | `/api/v1/project/generate-report/{id}` |


## Path Parameters

| Key       | Type      | Required     | Description                         |
| --------- | :-------: | :----------: | ----------------------------------- |
| id        | integer   | true         | a project ID to generate a report   |


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
url = /api/v1/project/generate-report/1
```

## Sample Response(s)
### A success Response
HTTP status 202 Accepted
```json
null
```

### An error response (case: a project ID doesn't exist of doesn't belong to a user ID in authorization token)
HTTP status 409 Conflict
```json
{
  "error": "a project ID doesn't exist of doesn't belong to a user ID"
}
```

### An error response (case: the project was already printed)
HTTP status 409 Conflict
```json
{
  "error": "the project was already printed"
}
```


### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
