# Golang Basics

## Overview
This Go project introduces backend development fundamentals with CRUD APIs for employee management. It leverages Gorilla/Mux for routing and GORM for ORM, aiming to provide a basic understanding of Go's backend capabilities.

## Technologies
- **Go**: Primary language.
- **Gorilla/Mux**: HTTP router.
- **GORM**: ORM library.

## Getting Started

### Prerequisites
- Go (1.19)
- SQL database (PostgreSQL)

### Installation
Clone and navigate:

```bash
git clone https://github.com/muf002/go-basic.git
cd go-basic
```

### Setup
Install dependencies:

```bash
go get -u
```

Configure `config/db.go` with your database details.

### Running
Start server:

```bash
go run main.go
```

## Endpoints
- Create: `POST /employee`
- Read: `GET /employee`, `GET /employee/{id}`
- Update: `PUT /employee/{id}`
- Delete: `DELETE /employee/{id}`
