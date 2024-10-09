# :writing_hand: <service-name>

Microservice designed to control 99minutos something

## Table of Contents

- [:writing\_hand: ](#writing_hand-)
    - [Table of Contents](#table-of-contents)
    - [Requirements](#requirements)
        - [Install Go](#install-go)
    - [Start](#start)
    - [Built with](#built-with)
    - [Structure](#structure)
    - [Infrastructure dependencies](#infrastructure-dependencies)
        - [Docker](#docker)
        - [Launch services](#launch-services)
        - [Rebuild app](#rebuild-app)
    - [Data Model](#data-model)
    - [Changelog](#changelog)

## Requirements

| Software | Version |
|:---------|:--------|
| Go       | 1.21.X  |

### Install Go

Windows

```shell
choco install golang
```

Debian based

```shell
apt install golang-go
```

Mac OS X

```shell
brew install go
```

Other installation methods: [Download Golang](https://go.dev/dl/)

## Start

1. Clone this project
2. Configure your environment variables ".env" file, you can use the ".env.example" file as a guide
3. Execute the following command to download the dependencies

    ```shell
    go mod tidy
    ```

4. For run the project execute the following command

    ```shell
    go run cmd/<service-name>/main.go
    ```

5. Happy hacking :D

## Start with docker

1. Install docker v24.0.2 [how to install?](https://docs.docker.com/engine/install/)
2. Install docker compose v2.18.1 [how to install?](https://docs.docker.com/compose/install/)
3. Clone this project
4. Copy environment configuration
    ```shell
    cp .env.example .env
    ```
5. Execute the project with the following command

    ```shell
    docker-compose up -d
    ```
6. Run mongodb seeders. Note: this command will execute the seeders in a default mongo db called `app`
    ```shell
    make mongodb-seeders
    ```
   if you want to customize the mongo db name (if you change MONGO_DATABASE in .env file), you can use the following
   command
    ```shell
    make mongodb-seeders MONGO_DATABASE=<your-db-name>
    ```
7. For next steps see [Trying the service (only for docker)](#trying-the-service-only-for-docker)

### Trying the service (only for docker)

with makefile

```shell
make testing-service
```

expected output

```bash
============================================================
Creating new order...

{"id":"6560ec6df49b452dade3e61e","first_name":"John","last_name":"Doe","sub_example":{"sub_example_id":123,"sub_example_name":"subExampleName"}}
============================================================
Searching order from seeder

{"id":"656045095ff16ef1a00fd4ef","first_name":"John","last_name":"Doe","sub_example":{"sub_example_id":123,"sub_example_name":"subExampleName"}}
============================================================
```

Or step by step

Create Example Order

```shell
curl --location --request POST '127.0.0.1:8080/api/v1/order/create'
```

Retrieve Example Order

```shell
curl --location '127.0.0.1:8080/api/v1/order/656045095ff16ef1a00fd4ef'
```

## Built with

- IntelliJ Ultimate with go plugin ( recommended )
- IntelliJ Goland
- Visual Studio Code

## Structure

```shell
<service-name>
.
├── CHANGELOG.md
├── CHANGELOG.template.md
├── Makefile
├── README.md
├── build
│   └── cloudbuild.yaml
├── cmd
│   └── service
│       └── main.go
├── docker
│   ├── Dockerfile.dev
│   └── Dockerfile
├── docker-compose.yml
├── docs
│   ├── MODEL.md
│   ├── STATE-MACHINE.md
│   └── TRANSITIONS.md
├── go.mod
├── go.sum
└── internal
    ├── domain
    │   ├── entities
    │   │   └── example.go
    │   ├── envs.go
    │   └── ports
    │       └── example_repo_iface.go
    ├── implementation
    │   └── example
    │       └── example_impl.go
    └── infrastructure
        ├── adapters
        │   └── repository
        │       └── mongo
        │           ├── order_impl.go
        │           └── seeders
        │               └── examples.json
        ├── driven
        │   ├── cmux
        │   │   └── cmux.go
        │   ├── core
        │   │   ├── envs.go
        │   │   └── logger.go
        │   ├── fiber_server
        │   │   └── fiber.go
        │   ├── mongodb
        │   │   └── mongodb.go
        │   ├── redis
        │   │   └── cache.go
        │   └── tracer
        │       └── tracer.go
        └── driver
            ├── grpc
            │   ├── domain_to_grpc.go
            │   ├── handlers.go
            │   └── server.go
            └── rest
                └── handlers.go


```

### Domain

The domain layer encapsulates the application's business logic, integrating data entities and ports that abstract this
logic.

### Implementation

The implementation layer is responsible for executing the specific technical details and interactions of an application.
It implements everything defined in the ports and utilizes the entities. Dependency injection is used in this layer to
facilitate the use of other services and repositories, thereby realizing the concepts defined in the domain layer
through concrete implementations and interfacing with external systems and managing data flow.

### Infrastructure

The architecture layer in software design provides the structural blueprint for the application. It outlines how various
components such as the domain, implementation, and presentation layers interact and are organized. This layer focuses on
aspects like scalability, maintainability, and technology stack, ensuring that the application's overall structure
supports its requirements and goals.

## Infrastructure dependencies

| Software | Version |
|:---------|:--------|
| MongoDB  | >=7.x   |
| Redis    | >=6.x   |

## Data Model

you can see the data model in the following link -> [Data Model](docs/MODEL.md)

## Changelog

you can see the changelog in the following link -> [Changelog](CHANGELOG.md)
