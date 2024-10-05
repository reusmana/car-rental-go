# Customer API Spec

## Create Customer API

Endpoint : POST http://localhost:8080/api/v1/customer

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

## Update Customer API

Endpoint : PUT http://localhost:8080/api/v1/customer/1

Request Body :

```json
{
  "name": "reusmana",
  "nik": "320206161198002",
  "phone": "085210780093"
}
```

Response Body Success :

```json
{
  "status": 200,
  "message": "customer updated successfully",
  "data": {
    "id": 1,
    "name": "reusmana",
    "nik": "320206161198002",
    "phone": "085210780093"
  }
}
```

## Get Customer API

Endpoint : GET (http://localhost:8080/api/v1/customer)/
:customerID

Response Body Success Without ID :

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

Response Body Success With ID :

```json
{
  "status": 200,
  "message": "success",
  "data": {
    "id": 1,
    "name": "reusmana",
    "nik": "320206161198002",
    "phone": "085210780093"
  }
}
```

## Remove Customer API

Endpoint : DELETE http://localhost:8080/api/v1/customer/1

Response Body Success :

```json
{
  "status": 200,
  "message": "customer deleted successfully"
}
```
