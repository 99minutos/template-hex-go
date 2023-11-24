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
3. Following the two previous steps from the earlier section
4. Execute the project with the following command

    ```shell
    docker-compose up -d
    ```
5. Run mongodb migrations

    ```shell
    make mongodb-seeders
    ```
    Note: aditionaly you can run the following command to see the logs
    ```shell    
      docker-compose logs -f
    ```

### Trying the service (only for docker)

with makefile

```shell
make testing-example-service:
```
expected output
```shell
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
├── build
│   ├── cloudbuild.yaml
│   └── Dockerfile
├── CHANGELOG.md
├── CHANGELOG.template.md
├── cmd
│   └── example
│       └── main.go
├── docker-compose.yml
├── Dockerfile.dev
├── docs
│   ├── MODEL.md
│   ├── STATE-MACHINE.md
│   └── TRANSITIONS.md
├── go.mod
├── go.sum
├── internal
│   ├── application
│   │   ├── ports
│   │   │   ├── example_repo_iface.go
│   │   │   └── example_service_iface.go
│   │   ├── repository
│   │   │   └── mongo
│   │   │       └── order_impl.go
│   │   └── services
│   │       └── example
│   │           ├── example_handlers_impl.go
│   │           └── example_impl.go
│   ├── config
│   │   ├── config.go
│   │   └── helpers.go
│   ├── domain
│   │   └── example.go
│   └── infraestructure
│       └── adapters
│           ├── driven
│           │   ├── cmux
│           │   │   └── cmux.go
│           │   ├── envs
│           │   │   └── config.go
│           │   ├── fiber
│           │   │   └── fiber.go
│           │   ├── logger
│           │   │   └── logger.go
│           │   ├── mongodb
│           │   │   └── mongodb.go
│           │   ├── redis
│           │   │   └── cache.go
│           │   └── tracer
│           │       └── tracer.go
│           └── driver
│               ├── grpc
│               │   ├── domain_to_grpc.go
│               │   ├── handlers.go
│               │   └── server.go
│               └── rest
│                   └── handlers.go
├── Makefile
└── README.md
```

## Infrastructure dependencies

| Software | Version |
|:---------|:--------|
| MongoDB  | >=5.x   |
| Redis    | >=6.x   |

## Data Model

you can see the data model in the following link -> [Data Model](docs/MODEL.md)

## Changelog

you can see the changelog in the following link -> [Changelog](CHANGELOG.md)
