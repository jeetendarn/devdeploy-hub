# ==========================================
# DevDeploy Hub Makefile
# ==========================================

APP_NAME=devdeploy-hub

# -----------------------------
# Go Commands
# -----------------------------

run:
	cd backend && go run cmd/server/main.go

build:
	cd backend && go build -o $(APP_NAME) cmd/server/main.go

clean:
	cd backend && rm -f $(APP_NAME)

fmt:
	cd backend && go fmt ./...

test:
	cd backend && go test ./...

deps:
	cd backend && go mod tidy

# -----------------------------
# Docker
# -----------------------------

docker-build-backend:
	docker build -t devdeploy-backend -f docker/backend/Dockerfile .

docker-build-frontend:
	docker build -t devdeploy-frontend -f docker/frontend/Dockerfile .

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

docker-restart:
	docker compose down
	docker compose up -d

# -----------------------------
# Kubernetes
# -----------------------------

k8s-apply:
	kubectl apply -f kubernetes/

k8s-delete:
	kubectl delete -f kubernetes/

k8s-pods:
	kubectl get pods -A

k8s-services:
	kubectl get svc -A

k8s-deployments:
	kubectl get deployments -A

# -----------------------------
# Minikube
# -----------------------------

minikube-start:
	minikube start

minikube-stop:
	minikube stop

minikube-dashboard:
	minikube dashboard

# -----------------------------
# Help
# -----------------------------

help:
	@echo ""
	@echo "========== DevDeploy Hub =========="
	@echo ""
	@echo "make run                 - Run backend"
	@echo "make build               - Build backend"
	@echo "make clean               - Remove binary"
	@echo "make fmt                 - Format Go code"
	@echo "make test                - Run tests"
	@echo "make deps                - Download dependencies"
	@echo ""
	@echo "make docker-build-backend"
	@echo "make docker-build-frontend"
	@echo "make docker-up"
	@echo "make docker-down"
	@echo "make docker-logs"
	@echo ""
	@echo "make k8s-apply"
	@echo "make k8s-delete"
	@echo "make k8s-pods"
	@echo "make k8s-services"
	@echo "make k8s-deployments"
	@echo ""
	@echo "make minikube-start"
	@echo "make minikube-stop"
	@echo "make minikube-dashboard"