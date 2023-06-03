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

| Key                 | Type       | Required  | Description                                 |
| ------------------- | :--------: | :-------: | ------------------------------------------- |
| Content-Type        | string     | true      | Content-Type has to be `application/json`   |


## Body Parameters

| Field Name      | Type     | Required | Default Value   |  Description                                                                                           |
| --------------- | -------- | -------- | --------------- | ------------------------------------------------------------------------------------------------------ |
| geolocation     | string   | true     |                 | latitude and longitude of location of an installed solar panel in (`latitude`,`longtitude`) format     |
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
