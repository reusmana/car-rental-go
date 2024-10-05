# Car API Spec

## Create Car API

Endpoint : POST http://localhost:8080/api/v1/cars

Request Body :

```json
{
  "brand": "honda",
  "model": "jazz",
  "daily_rent": 200000,
  "availability": true
}
```

Response Body Success :

```json
{
  "status": 201,
  "message": "success created cars",
  "data": {
    "id": 1,
    "brand": "honda",
    "model": "jazz",
    "daily_rent": 200000,
    "availability": true
  }
}
```

## Update Car API

Endpoint : PUT http://localhost:8080/api/v1/cars/1

Request Body :

```json
{
  "brand": "honda",
  "model": "jazz",
  "daily_rent": 400000,
  "availability": true
}
```

Response Body Success :

```json
{
  "status": 200,
  "message": "Car updated successfully",
  "data": {
    "id": 1,
    "brand": "honda",
    "model": "jazz",
    "daily_rent": 400000,
    "availability": true
  }
}
```

## Get Car API

Endpoint : GET (http://localhost:8080/api/v1/cars)/
:carID

Response Body Success Without ID :

```json
{
  "status": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "brand": "honda",
      "model": "jazz",
      "daily_rent": 400000,
      "availability": true
    }
  ]
}
```

Response Body Success With ID :

```json
{
  "status": 200,
  "message": "success",
  "data": {
    "id": 1,
    "brand": "honda",
    "model": "jazz",
    "daily_rent": 400000,
    "availability": true
  }
}
```

## Remove Car API

Endpoint : DELETE http://localhost:8080/api/v1/cars/1

Response Body Success :

```json
{
  "status": 200,
  "message": "Car deleted successfully"
}
```
