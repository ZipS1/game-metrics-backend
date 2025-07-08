![Game Metrics Logo](resources/logo.png)

# Game Metrics Backend

A microservices-based backend for board game score tracking service built with Go, featuring asynchronous communication through RabbitMQ. The system consists of five core microservices (API Gateway, Auth Service, Activities Service, Player Service, Game Service) designed to handle game scoring and player management efficiently.

**Frontend Application**: https://github.com/ZipS1/game-metrics-mobile-app

## 🚀 Quick Start

### Local Development (Docker Compose)

1. **Clone the repository:**

```bash
git clone https://github.com/ZipS1/game-metrics-backend.git
cd game-metrics-backend
```

2. **Start all services:**

```bash
docker-compose up -d
```

3. **Verify services are running:**

```bash
docker-compose ps
```


The API Gateway will be available at `http://localhost:8080`

### Kubernetes Deployment

1. **Apply the manifests:**

```bash
kubectl apply -f k8s/
```

2. **Check deployment status:**

```bash
kubectl get pods -n game-metrics
```

3. **Access the application:**

```bash
kubectl port-forward svc/api-gateway 8080:8080 -n game-metrics
```


## 🛠️ Technology Stack

### Core Technologies

- **Go 1.24** - Primary programming language
- **Gin Framework** - HTTP web framework for REST API
- **RabbitMQ** - Message broker for asynchronous communication
- **PostgreSQL** - Primary database with uuid-ossp extension
- **GORM** - ORM library for database operations


### Infrastructure \& Deployment

- **Docker** - Containerization platform
- **Docker Compose** - Local development orchestration
- **Kubernetes** - Production container orchestration
- **Nginx** - Reverse proxy and load balancer


### Development Tools

- **Zerolog** - Structured logging library
- **JWT** - Authentication and authorization
- **UUID** - Unique identifier generation


## 📐 Architecture Overview

The system follows a microservices architecture with the following services:

### Core Services

1. **API Gateway** ([`/api-gateway`](https://github.com/ZipS1/game-metrics-backend/tree/main/api-gateway))
    - Entry point for all client requests
    - Request routing and load balancing
    - Authentication middleware
2. **Auth Service** ([`/auth-service`](https://github.com/ZipS1/game-metrics-backend/tree/main/auth-service))
    - User authentication and authorization
    - JWT token management
    - User session handling
3. **Activities Service** ([`/activities-service`](https://github.com/ZipS1/game-metrics-backend/tree/main/activities-service))
    - Game activity tracking
    - Score calculation and processing
    - Activity history management
4. **Player Service** ([`/player-service`](https://github.com/ZipS1/game-metrics-backend/tree/main/player-service))
    - Player profile management
    - Player statistics and rankings
    - Social features
5. **Game Service** ([`/game-service`](https://github.com/ZipS1/game-metrics-backend/tree/main/game-service))
    - Game metadata management
    - Game rules and configuration
    - Game session handling

### Communication Pattern

The services communicate using the **choreography-based SAGA pattern** with event-driven architecture:

- **Synchronous**: HTTP/REST for direct service communication
- **Asynchronous**: RabbitMQ for event publishing and consuming
- **Idempotent**: All event handlers are designed to be idempotent

## ⚙️ Configuration

Each service’s example configuration file (`config.yml`) can be found in its `configs` directory:

- [api-gateway/configs/config.yml](https://github.com/ZipS1/game-metrics-backend/tree/main/api-gateway/configs/config.yml)
- [auth-service/configs/config.yml](https://github.com/ZipS1/game-metrics-backend/tree/main/auth-service/configs/config.yml)
- [activity-service/configs/config.yml](https://github.com/ZipS1/game-metrics-backend/tree/main/activity-service/configs/config.yml)
- [players-service/configs/config.yml](https://github.com/ZipS1/game-metrics-backend/tree/main/players-service/configs/config.yml)
- [game-service/configs/config.yml](https://github.com/ZipS1/game-metrics-backend/tree/main/game-service/configs/config.yml)

## 📚 API Documentation

An API Collection for testing with Bruno is available in the `docs/brunoApiCollections` [directory](https://github.com/ZipS1/game-metrics-backend/tree/main/docs/brunoApiCollections/dev) where you can import and run the collection in Bruno.
