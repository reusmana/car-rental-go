# Driver API Spec

## Create Driver API

Endpoint : POST http://localhost:8080/api/v1/driver

Request Body :

```json
{
  "name": "jong",
  "daily_cost": 50000,
  "incentive": 0
}
```

Response Body Success :

```json
{
  "status": 201,
  "message": "success created drivers",
  "data": {
    "id": 1,
    "name": "jong",
    "incentive": 0
  }
}
```

## Get Driver API

Endpoint : GET http://localhost:8080/api/v1/driver

Response Body Success :

```json
{
  "status": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "jong",
      "incentive": 0
    }
  ]
}
```
