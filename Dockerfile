## Simple multi-stage build (Debian runtime as requested)
FROM golang:1.22.2 AS builder
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Source
COPY . .

# Build (static)
ENV CGO_ENABLED=0
RUN go build -o gorgi ./src/main.go

FROM debian:12-slim
WORKDIR /app

# Copy binary only
COPY --from=builder /app/gorgi /usr/local/bin/gorgi

EXPOSE 8080

ENV GIN_MODE=release

ENTRYPOINT ["gorgi"]
CMD ["--host", "0.0.0.0", "--port", "8080"]
