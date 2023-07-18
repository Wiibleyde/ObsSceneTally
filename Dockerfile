# Builder stage
FROM golang:1.20.6 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o main main.go

# Final stage
FROM debian:latest

WORKDIR /root/

# Copy the built binary from the builder stage to the final stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/root/main"]
