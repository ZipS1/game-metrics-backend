SERVICES := api-gateway auth-service activity-service players-service game-service

BUILD_OPTIONS := -ldflags="-s -w"
BUILD_VARS := CGO_ENABLED=0

.PHONY: up down restart build clean vet lint tidy check

up:
	@docker compose up -d

down:
	@docker compose down -v --rmi local

restart: down up

build:
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		(cd $$service && \
		$(BUILD_VARS) go build $(BUILD_OPTIONS) -o ./target/$$service ./cmd/$$service/main.go;) \
	done

clean:
	@for service in $(SERVICES); do \
		echo "Cleaning $$service..."; \
		(cd $$service && rm -rf ./target;) \
	done

vet:
	@for service in $(SERVICES); do \
		echo "Vet $$service..."; \
		(cd $$service && go vet ./...;) \
	done

lint:
	@for service in $(SERVICES); do \
		echo "Lint $$service..."; \
		(cd $$service && golangci-lint run ./...;) \
	done

tidy:
	@for service in $(SERVICES); do \
		echo "Tidy $$service..."; \
		(cd $$service && go mod tidy;) \
	done

check: vet lint tidy
