# Genpasstore

Status: Educational project. Under development.
> ⚠️ Do not use in production. This project is for learning purposes only.

## Overview

CryptoStore is an experimental password management and encryption service written in Go.
It demonstrates concepts such as secure password storage, hashing, database integration, and a simple web interface.
The project is not audited, not secure, and should not be used with real data.

## Features

User registration and authentication

Password hashing and generation

RESTful API with JSON responses

Simple web UI built with static assets

Docker-based deployment

## Project Structure

```[]
password_store/
│
├── main.go                # Entry point of the application
├── db/db.go               # Database connection and queries
├── dto/user_dto.go        # User data transfer objects
├── handler/               # HTTP route handlers
│   ├── handler_password.go
│   └── handle_registry.go
├── password/              # Password utilities
│   ├── password_generate.go
│   └── password_hash.go
├── web/                   # Frontend assets (HTML, JS, CSS) build from frontend-crypto
│   ├── index.html
│   ├── favicon.ico
│   └── assets/
├── init.sql               # Database initialization script
├── Dockerfile             # Docker build configuration
├── docker-compose.yml     # Multi-container setup
├── .env                   # Environment variables
├── go.mod / go.sum        # Go dependencies
└── README.md
```

## Setup and Run

- Prerequisites
- Go 1.25+
- Docker and Docker Compose
- Run with Docker

## Build and start containers

`docker-compose up --build`

The service will start on <http://localhost:8080> (default port) or <http://your_ip_adress:8080>
Configuration variables are defined in the .env file.

Ensure your database is configured as expected in init.sql.

Important Notice

> This project is educational only.
> It has not been reviewed for security, reliability, or performance.
> Do not use it to store or process real user data, passwords, or any sensitive information.
> The authors are not responsible for any misuse or data loss.

## License

MIT License — for educational and research purposes only.
