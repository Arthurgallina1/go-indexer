FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# copy source code
COPY . .

RUN go build -o main ./api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Copy .env file if needed
COPY .env.docker .env

# Expose port
EXPOSE 3000

# Run the application
CMD ["./main"]