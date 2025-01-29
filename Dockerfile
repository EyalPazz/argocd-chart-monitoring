FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

# Copy the rest of the application code
COPY . .

RUN go build -o monitor cmd/monitor/main.go

# Stage 2: Create a smaller image with the Go binary
FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/monitor .

EXPOSE 8080

CMD ["./monitor"]
