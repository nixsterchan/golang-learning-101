# Simple Go API with MySQL + Docker

This project is a beginner-friendly example of building a simple REST API using **Go**, backed by a **MySQL** database. The application supports basic CRUD (Create, Read) operations on a list of items.

It uses Docker to containerize the Go app and the MySQL database, making setup and teardown easy and repeatable — perfect for learning or experimenting with Go-based web services.

- **Go** (`net/http`) to serve a simple API
- **MySQL** for backend storage
- **Docker** and **Docker Compose** for orchestration

The API supports basic operations to **add** and **list** "items" from a MySQL database.

---

## Project Components

```
.
├── Dockerfile              # Multi-stage Docker build for Go
├── docker-compose.yml      # Links MySQL and API services
├── main.go                 # Go application (REST API)
├── go.mod / go.sum         # Go module dependencies
├── .env                    # Configuration values for MySQL credentials and ports. [rmb to rename .env-template and fill it up]
├── .gitignore              # Files to be ignored.
├── .dockerignore           # Any files to be ignored by docker.

````

---

## Getting Started

### 1. Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

### 2. Spin Up the Project

To build and start the containers (ensure you are at the project root):

```bash
docker-compose up --build
````

This will:

* Build your Go binary in the first docker container
* Copies the main.exe to a second container that serves as the main app runne
* Start a MySQL database container
* Wait until MySQL is healthy
* Launch your API service on the specified port (default: `8080`)

---

### 3. Interact with the API

Once the API is running, use the following `curl` commands to test it.

#### Add an Item (POST)

```bash
curl -X POST http://localhost:8080/items \
     -H "Content-Type: application/json" \
     -d '{"name": "Banana"}'
```

#### Get All Items (GET)

```bash
curl http://localhost:8080/items
```

---

### 4. Tear It Down

When you're done, clean up the containers:

```bash
docker-compose down
```

If you want to also delete the MySQL volume (wipe the DB data):

```bash
docker-compose down -v
```

---

## Configuration

All configuration values are kept in the `.env` file (example):

```env
MYSQL_ROOT_PASSWORD=<choose-your-password>
MYSQL_DATABASE=<name-your-db>
MYSQL_PORT=3306
MYSQL_DSN=root:<choose-your-password>@tcp(mysql:3306)/<name-your-db>
API_PORT=8080
```

Make sure the values are consistent (especially `MYSQL_DSN`!).

---

## Stuff learnt

* How to build a really simple REST API with Go
* Working with MySQL and SQL drivers in Go
* Structuring JSON request/response handling
* Building multi-stage Docker images
* Managing services with Docker Compose
* Using environment variables for config

---

## Docker Commands Summary

| Command                     | Description                                |
| --------------------------- | ------------------------------------------ |
| `docker-compose up --build` | Build and start containers                 |
| `docker-compose down`       | Stop and remove containers                 |
| `docker-compose down -v`    | Stop and remove containers **and** volumes |
| `docker-compose ps`         | See running containers                     |
| `docker logs <container>`   | View logs from a container                 |

---

## 📚 References Used

* [Go Official Docs](https://golang.org/doc/)
* [Docker Compose Docs](https://docs.docker.com/compose/)
* [MySQL Go Driver](https://github.com/go-sql-driver/mysql)

