FROM golang:1.24.5-alpine AS builder
WORKDIR /src
RUN apk add --no-cache git
COPY my-go-backend/go.mod my-go-backend/go.sum ./
COPY restcountries-go-client /src/restcountries-go-client
COPY my-go-backend /src/my-go-backend
WORKDIR /src/my-go-backend
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go mod tidy
RUN go build -o /app/server .

FROM alpine:3.18
RUN apk add --no-cache ca-certificates curl

# Download ADOT collector binary
RUN curl -L https://github.com/aws-observability/aws-otel-collector/releases/latest/download/aws-otel-collector-linux-amd64 \
    -o /usr/local/bin/aws-otel-collector && \
    chmod +x /usr/local/bin/aws-otel-collector

# Copy Go binary
COPY --from=builder /app/server /usr/local/bin/server

# Copy ADOT config
COPY my-go-backend/config/adot-config.yaml /etc/adot/config.yaml

# Copy startup script
COPY my-go-backend/config/start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 8080

ENTRYPOINT ["/start.sh"]