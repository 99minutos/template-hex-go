# `<service-name>`

Microservice designed to control 99minutos something

## Table of Contents

- [Requirements](#requirements)
- [Getting started](#getting-started)
    - [Local (without Docker)](#local-without-docker)
    - [With Docker (recommended)](#with-docker-recommended)
- [Available commands](#available-commands)
- [API endpoints](#api-endpoints)
- [Built with](#built-with)
- [Structure](#structure)
- [Infrastructure dependencies](#infrastructure-dependencies)
- [Changelog](#changelog)

---

## Requirements

| Software       | Version |
|:---------------|:--------|
| Go             | 1.24.x  |
| Docker         | 24+     |
| Docker Compose | v2      |

### Install Go

**Windows**
```shell
choco install golang
```

**Debian / Ubuntu**
```shell
apt install golang-go
```

**macOS**
```shell
brew install go
```

Other methods: [Download Golang](https://go.dev/dl/)

---

## Getting started

### Local (without Docker)

1. Clone this project.
2. Copy and configure environment variables:
    ```shell
    cp .env.example .env
    ```
3. Download dependencies:
    ```shell
    make tidy
    ```
4. Run the service:
    ```shell
    make run
    ```

### With Docker (recommended)

1. Install [Docker](https://docs.docker.com/engine/install/) and [Docker Compose v2](https://docs.docker.com/compose/install/).
2. Clone this project.
3. Copy and configure environment variables:
    ```shell
    cp .env.example .env
    ```
4. Start all services (app + MongoDB + Redis):
    ```shell
    make up
    ```
5. To stop all containers:
    ```shell
    make down
    ```

---

## Available commands

Run `make help` to see the full list. Key targets:

| Command            | Description                                 |
|:-------------------|:--------------------------------------------|
| `make run`         | Run the app locally (outside container)     |
| `make fmt`         | Format all Go source files                  |
| `make tidy`        | Tidy and verify Go module dependencies      |
| `make vet`         | Run `go vet` on local source                |
| `make lint`        | Run `golangci-lint` (must be installed)     |
| `make test`        | Run tests locally with coverage             |
| `make coverage`    | Open HTML coverage report in browser        |
| `make up`          | Start all Docker services (with build)      |
| `make down`        | Stop and remove containers                  |
| `make build`       | Build Docker images                         |
| `make logs`        | Tail logs from all containers               |
| `make ps`          | List running containers                     |
| `make docker-test` | Run tests inside the container              |
| `make docker-vet`  | Run `go vet` inside the container           |

---

## API endpoints

Base URL: `http://localhost:8080`

### Health check

```shell
curl http://localhost:8080/api/v1/health
```

Expected response:

```json
{"status": "ok"}
```

### Create example order

```shell
curl -X POST http://localhost:8080/api/v1/order/create
```

### Get example order

```shell
curl http://localhost:8080/api/v1/order/<trackingId>
```

---

## Built with

- [Visual Studio Code](https://code.visualstudio.com/)
- [Zed](https://zed.dev/)
- [IntelliJ Ultimate with Go plugin](https://www.jetbrains.com/idea/)

---

## Structure

```
<service-name>
.
├── .dockerignore
├── .env.example
├── AGENTS.md
├── CHANGELOG.md
├── Makefile
├── build
│   └── cloudbuild.yaml
├── cmd
│   └── service
│       └── main.go
├── docker
│   ├── Dockerfile
│   └── Dockerfile.dev
├── docker-compose.yml
├── go.mod
├── go.sum
└── internal
    ├── domain
    │   ├── entities
    │   │   └── example.go
    │   ├── envs.go
    │   ├── errcodes
    │   │   └── errcodes.go
    │   ├── ports
    │   │   └── example.go
    │   ├── pubsub.go
    │   └── server
    │       ├── error.go
    │       └── pagination.go
    ├── helpers
    │   └── shortcodes
    │       └── validate.go
    ├── implementation
    │   └── example
    │       └── example_service.go
    └── infrastructure
        ├── adapters
        │   └── repository
        │       └── mongo
        │           ├── order.go
        │           └── seeders
        │               └── examples.json
        ├── driven
        │   ├── cmux
        │   │   └── cmux.go
        │   ├── core
        │   │   └── envs.go
        │   ├── dbg
        │   │   └── logger.go
        │   ├── fiber_server
        │   │   ├── fiber.go
        │   │   └── fiber_error.go
        │   ├── mongodb
        │   │   └── mongodb.go
        │   ├── redis
        │   │   └── cache.go
        │   ├── tracer
        │   │   └── tracer.go
        │   └── validation
        │       └── validate.go
        └── driver
            ├── grpc
            │   ├── domain_to_grpc.go
            │   ├── handlers.go
            │   └── server.go
            └── rest
                └── handlers.go
```

### Domain

The domain layer encapsulates the application's business logic, integrating data entities and ports that abstract this logic.

### Implementation

The implementation layer is responsible for executing the specific technical details and interactions of an application.
It implements everything defined in the ports and utilizes the entities. Dependency injection is used in this layer to
facilitate the use of other services and repositories, thereby realizing the concepts defined in the domain layer
through concrete implementations and interfacing with external systems and managing data flow.

### Infrastructure

The infrastructure layer provides the structural blueprint for the application. It outlines how the domain,
implementation, and presentation layers interact and are organized, focusing on scalability, maintainability, and
technology stack.

---

## Infrastructure dependencies

| Service | Image                 | Port  |
|:--------|:----------------------|:------|
| MongoDB | `mongo:7.0`           | 27017 |
| Redis   | `redis:6.2.13-alpine` | 6379  |

Both services are declared in `docker-compose.yml` and start automatically with `make up`.

---

## Changelog

See [CHANGELOG.md](CHANGELOG.md).
