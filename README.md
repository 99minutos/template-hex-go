# :writing_hand: <service-name>

Microservice designed to control 99minutos something

# Table of Contents

1. [Requirements](#requirements)
2. [Start](#start)
3. [Built with](#built-with)
4. [Structure](#structure)
5. [Data Model](#data-model)
6. [Changelog](#changelog)

## Requirements

It is necessary to install -> https://golang.org/

## Start

    1. Clone this project 
    2. Configure your environment variables ".env" file, you can use the ".env.example" file as a guide
    3. Execute the following command to download the dependencies
        $ go mod tidy
    4. For run the project execute the following command
        $ go run cmd/<service-name>/main.go
    5. Happy hacking :D

## Built with

    - IntelliJ Ultimate with go plugin ( recommended )
    - IntelliJ Goland
    - Visual Studio Code

## Structure

```
<service-name>
.
├── build
│ ├── cloudbuild.yaml
│ └── Dockerfile
├── CHANGELOG.md
├── CHANGELOG.template.md
├── cmd
│ └── example
│     └── main.go
├── docs
│ ├── MODEL.md
│ ├── STATE-MACHINE.md
│ └── TRANSITIONS.md
├── go.mod
├── go.sum
├── internal
│ ├── adapters
│ │ ├── ports
│ │ │ ├── example_repo_iface.go
│ │ │ └── example_service_iface.go
│ │ ├── repository
│ │ │ └── mongo
│ │ │     └── order_impl.go
│ │ └── services
│ │     └── example
│ │         ├── example_handlers_impl.go
│ │         └── example_impl.go
│ ├── config
│ │ ├── config.go
│ │ └── helpers.go
│ ├── domain
│ │ └── example.go
│ └── infraestructure
│     ├── driven
│     │ ├── cmux
│     │ │ └── cmux.go
│     │ ├── envs
│     │ │ └── config.go
│     │ ├── fiber
│     │ │ └── fiber.go
│     │ ├── logger
│     │ │ └── logger.go
│     │ ├── mongodb
│     │ │ └── mongodb.go
│     │ ├── redis
│     │ │ └── cache.go
│     │ └── tracer
│     │     └── tracer.go
│     └── driver
│         ├── grpc
│         │ ├── domain_to_grpc.go
│         │ ├── handlers.go
│         │ └── server.go
│         └── rest
│             └── handlers.go
├── logistics-source-of-truth.iml
├── Makefile
└── README.md



```

## Data Model

you can see the data model in the following link -> [Data Model](docs/MODEL.md)

## Changelog

you can see the changelog in the following link -> [Changelog](CHANGELOG.md)