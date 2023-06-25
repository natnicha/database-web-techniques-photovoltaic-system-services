# [GET] Get Solar Panel Models

To get an existing solar panel according to query parameters

## URL

| ** Method **    | GET                           | 
| --------------- | ----------------------------- | 
| ** Structure ** | `/api/v1/solar-panel-model`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                                                                                        |
| ------------------ | :-------: | :-------: | -------------------------------------------------------------------------------------------------- |
| filter             | array     | false     | to filter solar panel models for a specific condition e.g. id:1 means select only id = 1           |
| limit              | int       | false     | to limit number of solar panel models                                                              |
| offset             | int       | false     | to exclude from a response the first N items of a resource collection                              |
| sort_by            | string    | false     | to specify a sorting column in a resource collection e.g. id                                       |
| order_by           | string    | false     | to sort items specified in sort_by. Either `asc` for ascending sort or `desc` for descending sort  |


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
url = /api/v1/solar-panel-model?limit=2&offset=1&sort_by=id&order_by=asc
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
{
  "data": [
    {
      "id": 2,
      "name": "Jinko Solar - Tiger Neo 78HC-BDV",
      "description": "625 wp",
      "efficiency": "22.3600",
    },    
    {
      "id": 3,
      "name": "Jinko Solar - Tiger Pro 72HC Monofacial",
      "description": "540 wp",
      "efficiency": "21.3500",
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
