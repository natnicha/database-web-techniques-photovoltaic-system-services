# [POST] Create a New Product

To create a new product according to body parameters for a specific project and user in authorization token

## URL

| ** Method **    | POST                       | 
| --------------- | -------------------------- | 
| ** Structure ** | `/api/v1/product/create`   |


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

| Field Name           | Type     | Required | Default Value   |  Description                                                       |
| -------------------- | -------- | -------- | --------------- | ------------------------------------------------------------------ |
| project_id           | integer  | true     |                 | a project ID
| solar_panel_model_id | integer  | true     |                 | a solar panel model ID
| orientation          | string   | true     |                 | orientation of an installed solar panel Either N, E, S, W  |
| inclination          | numeric  | true     |                 | inclination or tilt of an installed solar panel in degree  |
| area                 | numeric  | true     |                 | area of an installed solar panel                           |
| latitude             | numeric  | true     |                 | latitude of location of an installed solar panel           |
| longitude            | numeric  | true     |                 | longitude of location of an installed solar panel          |


## Sample Request(s) 
```
url = /api/v1/project/create
```
```json
{
  "project_id": 1,
  "solar_panel_model_id": 1,
  "orientation": "N",
  "inclination": 5.25,
  "area": 20,
  "latitude": 50.8282,
  "longitude": 12.9209,
}
```

## Sample Response(s)
### A success Response
HTTP status 201 Created
```json
{
  "data": {
    "id": 1,
    "project_id": 1,
    "solar_panel_model_id": 1,
    "orientation": "N",
    "inclination": 5.25,
    "area": 20,
    "latitude": 50.8282,
    "longitude": 12.9209,
  }
}
```

### An error response (case: missing project ID)
HTTP status 400 Bad Request
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
