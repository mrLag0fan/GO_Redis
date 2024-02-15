# GO Redis and Postgres Project

This project provides Docker images for a PostgreSQL database and a Redis cache, along with Go code for interacting with them.

## Building the Docker Images

1. Ensure you have Docker installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the project directory in your terminal.

To build the Docker images, use the following commands:

### For PostgreSQL:

```bash
docker build -t go-postgres ./
```

### For Redis:

```bash
docker build -t redis-stack-server
```

These commands will build the Docker images using the respective Dockerfile provided in the repository.

## Running the Docker Containers

To run the Docker containers using the images you just built, use the following commands:

### For PostgreSQL:

```bash
docker run -d --name go-postgres -p 5432:5432 go-postgres
```

This command will start a new Docker container named `go-postgres` in detached mode (`-d` flag) and map port 5432 on your local machine to port 5432 in the container, allowing you to access the PostgreSQL database.

### For Redis:

```bash
docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
```

This command will start a new Docker container named `redis-stack-server` using the official Redis image `redis/redis-stack-server:latest` in detached mode (`-d` flag) and map port 6379 on your local machine to port 6379 in the container, allowing you to access the Redis cache.

## Project Structure

- `internal/entity`: Contains the `User` struct definition.
- `internal/postgres`: Contains functions for interacting with the PostgreSQL database.
- `internal/util`: Contains utility functions for converting between structs and maps.
- `internal/go_redis`: Contains functions for interacting with the Redis cache.
- `pkg/database`: Contains the database package, which initializes the PostgreSQL database and the Redis client.
- `pkg/error`: Contains the error package, which provides a function for checking and handling errors.

## How It Works

1. The `Post` function in `internal/go_redis` stores user data in both the Redis cache and the PostgreSQL database.
2. The `GetByUUID` function in `internal/go_redis` retrieves user data from the Redis cache if it exists, otherwise it falls back to fetching it from the PostgreSQL database.
3. The `StructTOMap` and `MapToUser` functions in `internal/util` are used for converting between the `User` struct and a map, which is required for interacting with Redis.
