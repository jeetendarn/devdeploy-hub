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
