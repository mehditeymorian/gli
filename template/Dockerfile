FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build main.go -o /main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /main .

ENTRYPOINT ["/main"]

