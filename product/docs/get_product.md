# [GET] Get Products

To get an existing products according to query parameters for a specific user in authorization token

## URL

| ** Method **    | GET                    | 
| --------------- | ---------------------- | 
| ** Structure ** | `/api/v1/product`      |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                                                                                        |
| ------------------ | :-------: | :-------: | -------------------------------------------------------------------------------------------------- |
| filter             | array     | false     | to filter products for a specific condition e.g. id:1 means select only id = 1                     |
| limit              | int       | false     | to limit number of products                                                                        |
| offset             | int       | false     | to exclude from a response the first N items of a resource collection                              |
| sort_by            | string    | false     | to specify a sorting column in a resource collection e.g. id                                       |
| order_by           | string    | false     | to sort items specified in sort_by. Either `asc` for ascending sort or `desc` for descending sort  |


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
url = /api/v1/product?limit=2&offset=1&sort_by=id&order_by=asc
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
{
  "data": [
    {
      "id": 1,
      "project_id": 1,
      "solar_panel_model_id": 1,
      "orientation": "N",
      "inclination": 5.25,
      "area": 20,
      "latitude": 50.8282,
      "longitude": 12.9209,
    },    
    {
      "id": 2,
      "project_id": 1,
      "solar_panel_model_id": 2,
      "orientation": "N",
      "inclination": 7.11,
      "area": 45.7,
      "latitude": 40.7128,
      "longitude": 74.0060,
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
