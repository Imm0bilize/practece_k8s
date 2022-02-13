FROM golang:1.17 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/.

FROM alpine:latest
WORKDIR /root/
COPY --from=builder ./app .
EXPOSE 8080
CMD ["./app"]
