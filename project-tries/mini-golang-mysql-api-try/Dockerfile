# Step 1: First docker container to perform the building of the go code
FROM golang:1.22.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.exe main.go

# Step 2: Set up the docker container that will be running the main executable
FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /app/main.exe .

# Install ca-certificates for https support
RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["./main.exe"]