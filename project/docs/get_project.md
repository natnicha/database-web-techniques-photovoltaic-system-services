# [GET] Get an Existing Project  

To get an existing project according to query parameters for a specific user in authorization token

## URL

| ** Method **    | POST                       | 
| --------------- | -------------------------- | 
| ** Structure ** | `/api/v1/project/update`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                                                                                                 |
| ------------------ | :-------: | :-------: | ----------------------------------------------------------------------------------------------------------- |
| filter             | array     | false     | to filter projects for a specific condition e.g. {"is_active": ["true"]} means select only is_active = true |
| limit              | int       | false     | to limit number of project for a specific user                                                              |
| offset             | int       | false     | to exclude from a response the first N items of a resource collection                                       |
| sort_by            | string    | false     | to specify a sorting column in a resource collection e.g. id, start_at                                      |
| order              | string    | false     | to sort items specified in sort_by. Either `asc` for ascending sort or `desc` for descending sort           |


## Header Parameters

| Key                 | Type       | Required  | Description                                 |
| ------------------- | :--------: | :-------: | ------------------------------------------- |
| Content-Type        | string     | true      | Content-Type has to be `application/json`   |
| Authorization       | string     | true      | A bearer token is required                  |


## Body Parameters

| Field Name   | Type     | Required | Default Value   |  Description               |
| ------------ | -------- | -------- | --------------- | -------------------------- |
|              |          |          |                 |                            |


## Sample Request(s) 
```
url = /api/v1/project/update?filter={"is_active": ["true"]}limit=1&offset=2&sort_by=id&order=asc
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
{
  "data": [
    {
      "id": 2,
      "name": "west Europe",
      "user_id": 1,
      "description": "located in west Europe",
      "star_at": "2023-05-30T00:00:01Z",
      "is_printed": true,
      "updated_at": "2023-06-01T00:00:00.00+02:00"
    },    
    {
      "id": 5,
      "name": "north Europe",
      "user_id": 1,
      "description": "located in north Europe",
      "star_at": "2023-05-30T00:00:01Z",
      "is_printed": true,
      "updated_at": "2023-06-01T00:00:00.00+02:00"
    }
  ]
}
```

### An error response (case: unsupported driver)
HTTP status 500 Internal Server Error
```json
{
  "error": "unsupported driver"
}
```