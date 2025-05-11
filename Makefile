SERVICES := api-gateway auth-service activity-service players-service game-service

BUILD_OPTIONS := -ldflags="-s -w"
BUILD_VARS := CGO_ENABLED=0

.PHONY: up down restart build clean vet lint tidy check k8s-build k8s-apply k8s-delete

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

k8s-build:
	@for service in $(SERVICES); do \
		echo "Building $${service} for k8s"; \
		(cd $$service && \
		docker build -f ./build/package/Dockerfile -t registry.gitlab.com/game-metrics/backend/$$service .. && \
		docker push registry.gitlab.com/game-metrics/backend/$$service) \
	done

k8s-apply:
	@for service in $(SERVICES) rabbitmq; do \
		echo "Running $${service} in k8s"; \
		(cd $$service && \
		kubectl apply -f ./deployments/k8s/dev/) \
	done

k8s-delete:
	@for service in $(SERVICES) rabbitmq; do \
		echo "Deleting $${service} in k8s"; \
		(cd $$service && \
		kubectl delete -f ./deployments/k8s/dev/) \
	done
