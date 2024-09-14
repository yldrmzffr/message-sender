# Message Sender Service

An automated message sending system built with Go.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
    - [Docker Deployment](#docker-deployment)
    - [Local Development](#local-development)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)

## Features

- Automated message sending at configurable intervals
- RESTful API endpoints for message management
- PostgreSQL integration for data persistence
- Redis for keeping provider ref. ids
- Swagger API documentation
- Docker support for easy deployment

## Requirements

- Go 1.23 or higher
- PostgreSQL 13
- Redis 6
- Docker and Docker Compose (optional)

## Installation

### Docker Deployment

1. Clone the repository:
   ```
   git clone https://github.com/yldrmzffr/message-sender.git
   cd message-sender
   ```

2. Start all services using Docker Compose:
   ```
   docker-compose up --build
   ```

   Note: When using Docker Compose, you don't need to create a `.env` file. The necessary environment variables are defined in the `docker-compose.yml` file.

3. The application will be accessible at `http://localhost:8080`.

### Local Development

1. Clone the repository:
   ```
   git clone https://github.com/yldrmzffr/message-sender.git
   cd message-sender
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up environment variables:
   Create a `.env` file in the project root and add the necessary variables (see [Configuration](#configuration) section).

4. Ensure PostgreSQL and Redis are running locally.

5. Start the application:
   ```
   go run cmd/api/main.go
   ```

## Configuration

| Environment Variable | Description                                | Default Value | Alternatives                                                      |
|----------------------|--------------------------------------------|---------------|-------------------------------------------------------------------|
| SERVICE_ENV | Application environment                    | dev | prod, staging                                                     |
| SERVICE_NAME | Application name                           | boilerplate | Any valid application name                                        |
| SERVICE_PORT | Application port                           | 8080 | Any valid port number                                             |
| LOG_LEVEL | Logging level                              | DEBUG | INFO, WARN, ERROR                                                 |
| DATABASE_HOST | PostgreSQL host                            | (Required) | localhost, postgres, db.example.com                              |
| DATABASE_PORT | PostgreSQL port                            | (Required) | 5432, Any valid port number                                      |
| DATABASE_USER | PostgreSQL username                        | (Required) | Any valid username                                                |
| DATABASE_PASSWORD | PostgreSQL password                        | (Required) | Any secure password                                               |
| DATABASE_DATABASE | PostgreSQL database name                   | (Required) | Any valid database name                                           |
| REDIS_URL | Redis URL                                  | (Required) | redis://localhost:6379/0, redis://redis:6380/0                   |
| NOTIFICATION_PROVIDER | Notification provider                      | mock | gcp                                                                |
| MESSAGES_AUTO_START | Auto-start message sending on app starting | true | false                                                              |
| MESSAGES_INTERVAL | Message sending interval (seconds)         | 120 | Any positive integer                                              |

To use these configurations for local development, you can set them as environment variables or include them in your `.env` file. For Docker deployments, they are already included in your `docker-compose.yml` file under the `environment` section of the relevant service.

Example `.env` file for local development:

```
SERVICE_ENV=dev
SERVICE_NAME=boilerplate
SERVICE_PORT=8080
LOG_LEVEL=DEBUG
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=user
DATABASE_PASSWORD=password
DATABASE_DATABASE=message_sender_db
REDIS_URL=redis://localhost:6379/0
NOTIFICATION_PROVIDER=mock
MESSAGES_AUTO_START=true
MESSAGES_INTERVAL=120
```

## API Documentation

Swagger UI is available at `http://localhost:8080/swagger/index.html` when the application is running.