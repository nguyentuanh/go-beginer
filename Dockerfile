FROM golang:1.19-alpine as builder
RUN apk add --no-cache --update gcc musl-dev g++ make git gnutls gnutls-dev gnutls-c++ bash git

WORKDIR /src

ADD ./go.mod ./go.sum ./
RUN go mod download

COPY cmd cmd
COPY config config
COPY internal internal
COPY pkg pkg

COPY .env.example .env.example

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o /dist/server cmd/server/*.go

FROM alpine:latest

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /dist/server /app/bin/server
COPY --from=builder /src/.env.example /app/bin/.env

WORKDIR /app/bin
EXPOSE 9000


