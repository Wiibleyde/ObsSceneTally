FROM golang:1.20.6 as builder

WORKDIR /app

COPY . .

RUN go build main.go

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/root/main"]
