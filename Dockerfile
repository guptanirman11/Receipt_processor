FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY server/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/receipt_processor_service

# Building nginx and running the service
FROM debian:latest

RUN apt-get update && apt-get install -y ca-certificates nginx && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/receipt_processor_service /root/

RUN chmod +x /root/receipt_processor_service

COPY client/ /usr/share/nginx/html/

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80 8080

CMD ["/bin/sh", "-c", "/root/receipt_processor_service & nginx -g 'daemon off;'"]
