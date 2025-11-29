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

## Setup and Run

- Prerequisites
- Go 1.25+
- Docker and Docker Compose
- Run with Docker

## Build and start containers

`docker-compose up --build`

The service will start on <http://localhost:8000> (default port) or <http://your_ip_adress:8000>
Configuration variables are defined in the .env file.

.env file look like 

```
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=password_manager
POSTGRES_PORT=5432
POSTGRES_HOST=localhost
SECRET_JWT=very_secret_token
```

Ensure your database is configured as expected in init.sql.

Important Notice

> This project is educational only.
> It has not been reviewed for security, reliability, or performance.
> Do not use it to store or process real user data, passwords, or any sensitive information.
> The authors are not responsible for any misuse or data loss.

## Build Postgresql separately

```bash
docker run -d \
  --name my-postgres \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=mydb \
  -v pgdata:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres:15
```


## License

MIT License — for educational and research purposes only.
