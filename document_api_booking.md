# Booking API Spec

## Create Booking API

Endpoint : POST http://localhost:8080/api/v1/booking

Request Body :

```json
{
  "customer_id": 1,
  "car_id": 1,
  "start_date": "2024-10-15",
  "end_date": "2024-10-16"
}
```

Response Body Success :

```json
{
  "status": 201,
  "message": "success created bookings",
  "data": {
    "id": 1,
    "customer_id": 1,
    "car_id": 1,
    "start_date": "2024-10-15",
    "end_date": "2024-10-16",
    "day_of_rent": 2,
    "total_cost": 400000,
    "status": true
  }
}
```

Response Body Error :

```json
{
  "status": 404,
  "message": "Car not available"
}
```

## Update Booking API

Endpoint : PUT http://localhost:8080/api/v1/customer/1

Request Body :

```json
{
  "customer_id": 1,
  "car_id": 2,
  "start_date": "2024-10-15",
  "end_date": "2024-10-15",
  "status": true
}
```

Response Body Success :

```json
{
  "status": 200,
  "message": "booking updated successfully",
  "data": {
    "id": 1,
    "customer_id": 1,
    "car_id": 2,
    "start_date": "2024-10-15",
    "end_date": "2024-10-15",
    "day_of_rent": 1,
    "total_cost": 300000,
    "status": true
  }
}
```

## Get Booking API

Endpoint : GET (http://localhost:8080/api/v1/booking)/
:bookingID

Response Body Success Without ID :

```json
{
  "status": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "customer_id": 1,
      "car_id": 1,
      "start_date": "2024-10-15",
      "end_date": "2024-10-16",
      "day_of_rent": 2,
      "total_cost": 400000,
      "status": true
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
    "customer_id": 1,
    "car_id": 1,
    "start_date": "2024-10-15",
    "end_date": "2024-10-16",
    "day_of_rent": 2,
    "total_cost": 400000,
    "status": true
  }
}
```

## Remove Booking API

Endpoint : DELETE http://localhost:8080/api/v1/customer/1

Response Body Success :

```json
{
  "status": 200,
  "message": "Booking deleted successfully"
}
```
