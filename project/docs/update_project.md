# [POST] Update an Existing Project  

To update an existing project according to body parameters for a specific user in authorization token

## URL

| ** Method **    | POST                       | 
| --------------- | -------------------------- | 
| ** Structure ** | `/api/v1/project/update`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


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

| Field Name   | Type     | Required | Default Value   |  Description                                                       |
| ------------ | -------- | -------- | --------------- | ------------------------------------------------------------------ |
| name         | string   | true     |                 | a project name                                                     |
| description  | string   | false    |                 | a project description                                              |
| star_at      | datetime | true     |                 | a starting project date in format YYYY-MM-DDThh:mm:dd e.g. 2023-05-30T00:00:01 |
| is_printed   | boolean  | false    | false           | In order to generate report, set to be `true` to export a report   |


## Sample Request(s) 
```
url = /api/v1/project/update
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
HTTP status 200 OK
```json
{
  "data": {
    "id": 1,
    "name": "Europe",
    "user_id": 1,
    "description": "located in central Europe",
    "star_at": "2023-05-30T00:00:01Z",
    "is_printed": false,
    "updated_at": "2023-06-01T00:00:00.00+02:00"
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
