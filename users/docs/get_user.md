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

| Key                 | Type       | Required  | Description                    |
| ------------------- | :--------: | :-------: | ------------------------------ |
|                     |            |           |                                |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description      |
| ---------- | ------- | -------- | --------------- | ----------------- |
|            |         |          |                 |                   |


## Sample Request(s) 

### A sample request where user ID = 1
```
url = /api/v1/user/1
```

## Sample Response(s) 
### A success response
```json
{
  "status" : "SUCCESS",
  "code" : 200,
  "data": {
    "id" : "1",
    "first_name" : "Gloria",
    "last_name" : "Bonner",
    "email" : "gloria.bonner@gmail.com",
    "is_active" : "true"
  }
}
```
