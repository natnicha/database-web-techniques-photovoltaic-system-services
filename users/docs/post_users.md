# [POST] Create a New User  

To create a new user according to body parameters

## URL

| ** Method **    | POST                    | 
| --------------- | ----------------------- | 
| ** Structure ** | `/api/v1/user/create`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                    |
| ------------------- | :--------: | :-------: | ------------------------------ |
|                     |            |           |                                |


## Body Parameters

| Field Name | Type    | Required | Default Value   |  Description                                                                          |
| ---------- | ------- | -------- | --------------- | ------------------------------------------------------------------------------------- |
| first_name | string  | true     |                 | a user's first name                                                                   |
| last_name  | string  | true     |                 | a user's last name                                                                    |
| email      | string  | true     |                 | a user's email using for login to application                                         |
| is_active  | boolean | false    | true            | In order to activate this new user, set to be `true` to activate. Otherwise, `false`  |

## Sample Request(s) 
```
url = /api/v1/user/create
```
```json
{
  "first_name" : "Natnicha",
  "last_name" : "Rodtong",
  "email" : "nat.rodtong@gmail.com",
  "is_active" : "true"
}
```

## Sample Response(s)
### Success Response
```json
{
  "status" : "SUCCESS",
  "code" : 201,
  "data": {
    "id" : "1",
    "first_name" : "Natnicha",
    "last_name" : "Rodtong",
    "email" : "nat.rodtong@gmail.com",
    "is_active" : "true"
  }
}
```

### Error Response (case: missing first name)
```json
{
  "status" : "ERROR",
  "code" : 400,
  "message" : "Bad Request"
}
```
