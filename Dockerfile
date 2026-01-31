# Builder Stage
FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker-registry-exporter

# Final Image
FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/docker-registry-exporter .

COPY entrypoint.sh /app/entrypoint.sh

# Set default environment variable for the registry address
ENV REGISTRY_ADDRESS=http://docker.io

# Install curl for health check and make the entrypoint script executable
RUN apk add --no-cache curl && \
    chmod +x /app/entrypoint.sh && \
    # Add a non-root user for better security
    adduser -D exporteruser && \
    chown -R exporteruser:exporteruser /app

USER exporteruser

EXPOSE 9055

HEALTHCHECK --interval=30s --timeout=20s --start-period=120s --retries=3 \
  CMD curl --fail http://localhost:9055/metrics || exit 1

ENTRYPOINT ["/app/entrypoint.sh"]