FROM golang:1.24.5-alpine AS builder
WORKDIR /src
RUN apk add --no-cache git

# copy backend module files (from repo root)
COPY my-go-backend/go.mod my-go-backend/go.sum ./
# copy local client module
COPY restcountries-go-client /src/restcountries-go-client
# copy backend source
COPY my-go-backend /src/my-go-backend

WORKDIR /src/my-go-backend
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go mod tidy
RUN go build -o /app/server .

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /usr/local/bin/server
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/server"]