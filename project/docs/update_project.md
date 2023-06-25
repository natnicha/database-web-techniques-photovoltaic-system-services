# [PUT] Update an Existing Project  

To update an existing project according to body parameters for a specific user in authorization token

## URL

| ** Method **    | PUT                            | 
| --------------- | ------------------------------ | 
| ** Structure ** | `/api/v1/project/update/{id}`  |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
| id        | integer   | true         | a project ID to update          |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                                                                   | Permission         |
| ------------------- | :--------: | :-------: | ----------------------------------------------------------------------------- | ------------------ |
| Content-Type        | string     | true      | Content-Type has to be `application/json`                                     | internal, external |
| Authorization       | string     | false     | A bearer token is required for external access                                | external only      |
| api-key             | string     | false     | Instead of Authorization, internal requests required only api-key and user-id | internal only      |
| user-id             | string     | false     | Instead of Authorization, internal requests required only api-key and user-id | internal only      |


## Body Parameters

| Field Name   | Type     | Required | Default Value   |  Description                                                       |
| ------------ | -------- | -------- | --------------- | ------------------------------------------------------------------ |
| name         | string   | true     |                 | a project name                                                     |
| description  | string   | false    | `blank`         | a project description                                              |
| start_at     | datetime | true     |                 | a starting project date in format YYYY-MM-DDThh:mm:dd e.g. 2023-05-30T00:00:01 |
| is_printed   | boolean  | false    | false           | In order to generate report, set to be `true` to export a report   |


## Sample Request(s) 
```
url = /api/v1/project/update/1
```
```json
{
    "name": "Europe",
    "description": "located in central Europe",
    "is_printed": false,
    "start_at": "2023-05-30T00:00:01Z"
}
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
{
  "data": {
    "id": 1,
    "name": "Europe",
    "user_id": 1,
    "description": "located in central Europe",
    "start_at": "2023-05-30T00:00:01Z",
    "is_printed": false,
  }
}
```

### An error response (case: missing project name)
HTTP status 400 Bad Request
```json
null
```

### An error response (case: a project ID doesn't belong to user ID in authorization token)
HTTP status 409 Conflict
```json
{
  "error": "a project ID doesn't belong to a user ID"
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```
