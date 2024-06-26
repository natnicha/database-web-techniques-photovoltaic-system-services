# [POST] Create a New Project  

To create a new project according to body parameters for a specific user

## URL

| ** Method **    | POST                       | 
| --------------- | -------------------------- | 
| ** Structure ** | `/api/v1/project/create`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


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
| description  | string   | false    |                 | a project description                                              |
| star_at      | datetime | true     |                 | a starting project date in format YYYY-MM-DDThh:mm:dd e.g. 2023-05-30T00:00:01 |
| is_printed   | boolean  | false    | false           | In order to generate report, set to be `true` to export a report   |


## Sample Request(s) 
```
url = /api/v1/project/create
```
```json
{
  "name": "Europe",
  "description": "located in central Europe",
  "is_printed": false,
  "star_at": "2023-05-30T00:00:01Z"
}
```

## Sample Response(s)
### A success Response
HTTP status 201 Created
```json
{
  "data": {
    "id": 1,
    "name": "Europe",
    "user_id": 1,
    "description": "located in central Europe",
    "star_at": "2023-05-30T00:00:01Z",
    "is_printed": false
  }
}
```

### An error response (case: missing project name)
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
