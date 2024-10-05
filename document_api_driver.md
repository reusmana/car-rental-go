# Driver API Spec

## Create Driver API

Endpoint : POST http://localhost:8080/api/v1/driver

Request Body :

```json
{
  "name": "resuamana",
  "nik": "320206161198002",
  "phone": "085210780093"
}
```

Response Body Success :

```json
{
  "status": 201,
  "message": "success created customers",
  "data": {
    "id": 1,
    "name": "resuamana",
    "nik": "320206161198002",
    "phone": "085210780093"
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
      "name": "reusmana",
      "nik": "320206161198002",
      "phone": "085210780093"
    }
  ]
}
```
