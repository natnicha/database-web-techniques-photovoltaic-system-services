# [POST] Gather All Locations for Weather Scraping

To gather all locations before sending to a target API

## URL

| ** Method **    | POST                      | 
| --------------- | ------------------------- | 
| ** Structure ** | `/api/v1/weather/daily`   |


## Path Parameters

| Key       | Type      | Required     | Description                     |
| --------- | :-------: | :----------: | ------------------------------- |
|           |           |              |                                 |


## Query Parameters

| Key                 | Type       | Required  | Description                                                       | Permission         |
| ------------------- | :--------: | :-------: | ----------------------------------------------------------------- | ------------------ |
| api-key             | string     | false     | Instead of Authorization, internal requests required only api-key | internal only      |


## Header Parameters

| Key                 | Type       | Required  | Description                                 |
| ------------------- | :--------: | :-------: | ------------------------------------------- |
|                     |            |           |                                             |


## Body Parameters

| Field Name      | Type     | Required | Default Value   |  Description       |
| --------------- | -------- | -------- | --------------- | ------------------ |
|                 |          |          |                 |                    |


## Sample Request(s) 
```
url = /api/v1/weather/daily
```

## Sample Response(s)
### A success Response
HTTP status 202 Accepted
```json
null
```
