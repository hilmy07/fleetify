# Fleetify Backend API

Fleetify is a vehicle maintenance management backend built with Go, Fiber, GORM, MySQL, and Docker.

This project provides APIs for:

* Vehicle maintenance reporting
* Maintenance approval workflow
* Spare part and service item management
* Maintenance completion process

---

# Tech Stack

* Go
* Fiber
* GORM
* MySQL
* Docker
* Docker Compose

---

# Project Structure

```bash
fleetify/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   ├── controllers/
│   ├── dto/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── seeder/
│   └── services/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── .env
```

---

# Features

## Authentication

* Login endpoint
* Simple role-based access

## Roles

* SA (Service Advisor)
* APPROVAL

## Maintenance Report Workflow

### Service Advisor

* Create maintenance report
* Add maintenance items
* Upload initial vehicle photo

### Approval

* Approve maintenance report

### Completion

* Complete maintenance report
* Upload proof photo

---

# Database Schema

## Users

| Field    | Type   |
| -------- | ------ |
| id       | uint   |
| username | string |
| password | string |
| role     | string |

---

## Vehicles

| Field         | Type   |
| ------------- | ------ |
| id            | uint   |
| license_plate | string |
| model         | string |

---

## Master Items

| Field     | Type   |
| --------- | ------ |
| id        | uint   |
| item_name | string |
| type      | string |
| price     | float  |

---

## Maintenance Reports

| Field        | Type   |
| ------------ | ------ |
| id           | uint   |
| vehicle_id   | uint   |
| created_by   | uint   |
| odometer     | uint   |
| complaint    | string |
| status       | string |
| initial_post | string |
| proof_photo  | string |
| created_at   | string |

---

## Report Items

| Field          | Type   |
| -------------- | ------ |
| id             | uint   |
| report_id      | uint   |
| item_id        | uint   |
| quantity       | uint   |
| notes          | string |
| price_snapshot | float  |

---

# Installation

## Clone Repository

```bash
git clone https://github.com/hilmy07/fleetify.git
cd fleetify
```

---

# Environment Variables

Create `.env` file:

```env
PORT=5000

DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=fleetify
```

---

# Run With Docker

```bash
docker compose up --build
```

Application will run at:

```bash
http://localhost:5000
```

---

# Seeded Data

## Users

| Username         | Role     |
| ---------------- | -------- |
| service_advisor  | SA       |
| approval_manager | APPROVAL |

---

## Vehicles

| ID | License Plate | Model              |
| -- | ------------- | ------------------ |
| 1  | B 1234 ABC    | Toyota Avanza      |
| 2  | L 5678 DEF    | Honda Brio         |
| 3  | N 9012 GHI    | Mitsubishi Xpander |

---

## Master Items

| ID | Item            | Type    | Price  |
| -- | --------------- | ------- | ------ |
| 1  | Engine Oil      | PART    | 250000 |
| 2  | Oil Filter      | PART    | 75000  |
| 3  | Brake Pad       | PART    | 350000 |
| 4  | General Service | SERVICE | 500000 |
| 5  | Wheel Alignment | SERVICE | 200000 |

---

# API Endpoints

## Health Check

### GET /

Response:

```json
{
  "message": "Backend is running well"
}
```

---

# Reports

## Create Maintenance Report

### POST /api/reports

Request:

```json
{
  "vehicle_id": 2,
  "odometer": 125000,
  "complaint": "Engine produces abnormal noise when accelerating",
  "initial_photo": "https://example.com/uploads/report-photo.jpg",
  "items": [
    {
      "item_id": 3,
      "quantity": 2,
    },
    {
      "item_id": 1,
      "quantity": 1,
    }
  ]
}
```

Response:

```json
{
  "message": "report created successfully"
}
```

---

## Get All Reports

### GET /api/reports

---

## Approve Report

### PATCH /api/reports/:id/approve

Example:

```bash
PATCH /api/reports/1/approve
```

Response:

```json
{
  "message": "report approved successfully"
}
```

---

## Complete Report

### PATCH /api/reports/:id/complete

Request:

```json
{
  "proof_photo": "https://example.com/uploads/completed-photo.jpg"
}
```

Response:

```json
{
  "message": "report completed successfully"
}
```

---

# Docker Commands

## Build Container

```bash
docker compose build
```

## Run Container

```bash
docker compose up
```

## Run in Detached Mode

```bash
docker compose up -d
```

## Stop Container

```bash
docker compose down
```

---

# Access MySQL Container

```bash
docker exec -it fleetify_mysql mysql -u root -p
```

Use database:

```sql
USE fleetify;
```

Check tables:

```sql
SHOW TABLES;
```

Check data:

```sql
SELECT * FROM vehicles;
SELECT * FROM maintenance_reports;
SELECT * FROM report_items;
```

---

# Example Workflow

1. Service Advisor creates maintenance report
2. Approval manager approves report
3. Technician completes maintenance
4. System stores maintenance history

---

# Future Improvements

* JWT Authentication
* Upload image storage
* Pagination
* Swagger Documentation
* Unit Testing
* Role Middleware
* Soft Delete
* Audit Log
* Dashboard Analytics

---

# Author

Hilmy Haidar

GitHub:

[https://github.com/hilmy07](https://github.com/hilmy07)
