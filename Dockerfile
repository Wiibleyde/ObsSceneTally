# Builder stage
FROM golang:1.20.6 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o main main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage to the final stage
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

RUN mkdir -p /root/data
RUN mkdir -p /root/logs

EXPOSE 8080

CMD ["/root/main"]
