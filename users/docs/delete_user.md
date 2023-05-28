# [DELTE] Delete an Existing User  

To delete an existing user by user ID

## URL

| ** Method **    | DELETE                       | 
| --------------- | ---------------------------- | 
| ** Structure ** | `/api/v1/user/delete/{id}`   |


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

| Field Name | Type    | Required | Default Value   |  Description      |
| ---------- | ------- | -------- | --------------- | ----------------- |
|            |         |          |                 |                   |


## Sample Request(s) 
```
url = /api/v1/user/delete/1
```

## Sample Response(s)
### A success Response
```json
{
  "status" : "SUCCESS",
  "code" : 200
}
```

### An error response (case: user ID doesn't exist)
```json
{
  "status" : "ERROR",
  "code" : 400,
  "message" : "Bad Request"
}
```
