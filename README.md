# DevDeploy Hub

A lightweight DevOps portfolio project demonstrating modern backend development and DevOps practices.

## Tech Stack

### Frontend
- HTML5
- CSS3
- JavaScript (ES6)

### Backend
- Go
- Gin Framework

### Database
- PostgreSQL

### DevOps (Planned)
- Docker
- Docker Compose
- Nginx
- Kubernetes
- GitHub Actions
- Prometheus
- Grafana
- Loki
- Terraform

---

## Current Features

- Project structure
- Gin HTTP server
- Environment configuration
- PostgreSQL connectivity
- Health API
- Projects API
- Modern responsive frontend

---
cd frontend

live-server
---
## Run Backend

```bash
cd backend

go run cmd/server/main.go

```

Server:

```
http://localhost:8080
```

Health API:

```
GET /api/health
```

Projects API:

```
GET /api/projects
```

---

## Project Structure

```
frontend/
backend/
database/
docker/
kubernetes/
monitoring/
terraform/
docs/
```

---

## Git History

- Initial project structure
- Frontend UI
- Backend configuration
- PostgreSQL integration
- Project API

## Docker - Backend

### Features

- Multi-stage Docker build
- Alpine Linux runtime
- Small production image
- Go binary containerization
- Port 8080 exposed

### Build Image

```bash
docker build -t devdeploy-backend -f docker/backend/Dockerfile .
```

### Run Container

```bash
docker run --rm -p 8080:8080 devdeploy-backend
```

### Useful Commands

```bash
docker images
docker ps
docker logs <container_id>
docker exec -it <container_id> sh
docker stop <container_id>
```

## Docker - Frontend

### Base Image

- Nginx Alpine

### Features

- Static file hosting
- Custom Nginx configuration
- Lightweight production image

### Build

```bash
docker build -t devdeploy-frontend -f docker/frontend/Dockerfile .
```

### Run

```bash
docker run --rm -p 3000:80 devdeploy-frontend
```

### Test

Open:

- http://localhost:3000


## Docker Compose

### Services

- Frontend (Nginx)
- Backend (Go)
- PostgreSQL
- Redis

### Build

```bash
docker compose build
```

### Start

```bash
docker compose up -d
```

### Stop

```bash
docker compose down
```

### View Containers

```bash
docker ps
```

### View Logs

```bash
docker logs devdeploy-backend
docker logs devdeploy-postgres
docker logs devdeploy-frontend
docker logs devdeploy-redis
```

### Project URLs

Frontend:
http://localhost:3000

Backend:
http://localhost:8080/api/health

Projects API:

http://localhost:8080/api/projects