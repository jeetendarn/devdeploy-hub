# ===============================
# Stage 1 - Build
# ===============================

FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o devdeploy-hub ./cmd/server

# ===============================
# Stage 2 - Runtime
# ===============================

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/devdeploy-hub .

COPY backend/.env .

EXPOSE 8080

CMD ["./devdeploy-hub"]