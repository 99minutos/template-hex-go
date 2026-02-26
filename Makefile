APP_NAME=app
MAIN=cmd/service/main.go
COVERAGE_OUT=coverage.out

.PHONY: help \
        run fmt tidy vet lint test coverage \
        up down build logs ps \
        docker-up docker-down docker-build docker-logs docker-ps

# ── Help ──────────────────────────────────────────────────────────────────────

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

# ── Local development ─────────────────────────────────────────────────────────

run: ## Run the app locally (outside container)
	go run $(MAIN)

fmt: ## Format all Go source files
	go fmt ./...

tidy: ## Tidy and verify Go module dependencies
	go mod tidy && go mod verify

vet: ## Run go vet on local source
	go vet ./...

lint: ## Run golangci-lint (must be installed)
	golangci-lint run ./...

test: ## Run tests locally with coverage
	go test ./... -v -coverprofile=$(COVERAGE_OUT)

coverage: test ## Open HTML coverage report in browser
	go tool cover -html=$(COVERAGE_OUT)

# ── Docker shortcuts ──────────────────────────────────────────────────────────

up: docker-up ## Alias for docker-up

down: docker-down ## Alias for docker-down

build: docker-build ## Alias for docker-build

logs: docker-logs ## Alias for docker-logs

ps: docker-ps ## Alias for docker-ps

# ── Docker commands ───────────────────────────────────────────────────────────

docker-up: ## Start all services (with build)
	docker compose up --build

docker-down: ## Stop and remove containers
	docker compose down

docker-build: ## Build Docker images
	docker compose build

docker-logs: ## Tail logs from all containers
	docker compose logs -f

docker-ps: ## List running containers
	docker compose ps

docker-test: ## Run tests inside the container
	docker compose run --rm $(APP_NAME) go test ./... -v -coverprofile=$(COVERAGE_OUT)

docker-vet: ## Run go vet inside the container
	docker compose run --rm $(APP_NAME) go vet ./...

# -- IA commands (if applicable) ─────────────────────────────────────────────────────

link-gemini: ## Link AGENTS.md to Gemini (if applicable)
	ln -s AGENTS.md GEMINI.md

link-claude: ## Link AGENTS.md to Claude (if applicable)
	ln -s AGENTS.md CLAUDE.md
