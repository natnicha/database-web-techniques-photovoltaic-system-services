# [GET] Get an Existing Project  

To get an existing project according to query parameters for a specific user in authorization token

## URL

| ** Method **    | GET                  | 
| --------------- | -------------------- | 
| ** Structure ** | `/api/v1/project`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                                                                                                 |
| ------------------ | :-------: | :-------: | ----------------------------------------------------------------------------------------------------------- |
| filter             | array     | false     | to filter projects for a specific condition e.g. id:1 means select only id = 1. Over 1 filter values are separated by comma e.g. id:1,project_id:1 
| limit              | int       | false     | to limit number of project for a specific user                                                              |
| offset             | int       | false     | to exclude from a response the first N items of a resource collection                                       |
| sort_by            | string    | false     | to specify a sorting column in a resource collection e.g. id, start_at                                      |
| order_by           | string    | false     | to sort items specified in sort_by. Either `asc` for ascending sort or `desc` for descending sort           |


## Header Parameters

| Key                 | Type       | Required  | Description                                                                   | Permission         |
| ------------------- | :--------: | :-------: | ----------------------------------------------------------------------------- | ------------------ |
| Authorization       | string     | false     | A bearer token is required for external access                                | external only      |
| api-key             | string     | false     | Instead of Authorization, internal requests required only api-key and user-id | internal only      |
| user-id             | string     | false     | Instead of Authorization, internal requests required only api-key and user-id | internal only      |


## Body Parameters

| Field Name   | Type     | Required | Default Value   |  Description               |
| ------------ | -------- | -------- | --------------- | -------------------------- |
|              |          |          |                 |                            |


## Sample Request(s) 
```
url = /api/v1/project?filter=is_printed:true&limit=2&offset=1&sort_by=id&order_by=asc
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
    },    
    {
      "id": 5,
      "name": "north Europe",
      "user_id": 1,
      "description": "located in north Europe",
      "star_at": "2023-05-30T00:00:01Z",
      "is_printed": true,
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
