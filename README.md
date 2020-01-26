## Abstract
An API for the BusSaver mobile app. Wraps around LTA's Bus Services API but only returns required values to the mobile front-end.

## Usage

### Get metadata of the incoming buses of a bus service

**Definition**

`GET /incoming-buses/{busStopCode}/{serviceNumber}`

**Example: Get metadata of incoming buses of the bus 4 service at "Bef Changi Prison Cplx"**

`https://bus-saver-server.herokuapp.com/incoming-buses/97331/4`

**Response**

- `200 OK`

```json
[
   {
      "estimatedArrival":"2020-01-26T09:09:33+08:00",
      "load":"SEA",
      "isWheelChairAccessible":true,
      "type":"SD"
   },
   {
      "estimatedArrival":"2020-01-26T09:24:27+08:00",
      "load":"SEA",
      "isWheelChairAccessible":true,
      "type":"SD"
   },
   {
      "estimatedArrival":"2020-01-26T09:37:27+08:00",
      "load":"SEA",
      "isWheelChairAccessible":true,
      "type":"SD"
   }
]
```

## Deploying on Heroku

Guide: https://leanpub.com/howtodeployagowebapptoheroku101/read