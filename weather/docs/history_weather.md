# [POST] Gather historical weather for a specific location

To gather historical weather for a specific location according request body parameters

## URL

| ** Method **    | POST                      | 
| --------------- | ------------------------- | 
| ** Structure ** | `/api/v1/weather/history` |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                | Type      | Required  | Description                   |
| ------------------ | :-------: | :-------: | ----------------------------- |
|                    |           |           |                               |


## Header Parameters

| Key                 | Type       | Required  | Description                                                       | Permission         |
| ------------------- | :--------: | :-------: | ----------------------------------------------------------------- | ------------------ |
| Content-Type        | string     | true      | Content-Type has to be `application/json`                         | internal, external |
| api-key             | string     | false     | Instead of Authorization, internal requests required only api-key | internal only      |


## Body Parameters

| Field Name      | Type     | Required | Default Value   |  Description                                                                                           |
| --------------- | -------- | -------- | --------------- | ------------------------------------------------------------------------------------------------------ |
| latitude        | string   | true     |                 | latitude of location of an installed solar panel                                                       |
| longitude       | string   | true     |                 | longitude of location of an installed solar panel                                                      |
| start_at        | datetime | true     |                 | a starting expected date for historical weather in YYYY-MM-DDThh:mm:dd format e.g. 2023-05-30T00:00:01 |
| end_at          | datetime | true     |                 | a ending expected date for historical weather in YYYY-MM-DDThh:mm:dd format e.g. 2023-06-30T00:00:01   |


## Sample Request(s) 
```
url = /api/v1/weather/history
```

## Sample Response(s)
### A success Response
HTTP status 200 OK
```json
null
```
