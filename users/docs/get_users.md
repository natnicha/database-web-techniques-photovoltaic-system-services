# [GET] Get User Information 

To get user information such as first name, last name, email, etc.

## URL

| ** Method **    | GET                       | 
| --------------- | ------------------------- | 
| ** Structure ** | `/api/v1/user/{id}`       |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
| id        | integer   | true         | user ID                         |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                                   |
| ------------------- | :--------: | :-------: | --------------------------------------------- |
| Content-Type        | string     | true      | Content-Type has to be `application/json`     |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description      |
| ---------- | ------- | -------- | --------------- | ----------------- |
|            |         |          |                 |                   |


## Sample request(s) 

### Sample Query Parameter Requested 
```
url = /api/v1/user/1
```

### Success Response
```json
{
  "status" : "SUCCESS",
  "code" : 200,
  "data": {
    "id" : "1",
    "first_name" : "Natnicha",
    "last_name" : "Rodtong",
    "email" : "nat.rodtong@gmail.com",
    "is_active" : "true"
  }
}
```

## Error Response (case: bad request)
```json
{
  "status" : "ERROR",
  "code" : 400,
  "message" : "Bad Request"
}
```
