FROM golang:1.16-alpine AS builder
LABEL maintainer="Mayshev Alex <alex.mayshev.86@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN GOOS=linux go build -v -o ./bin/bot ./cmd/bot

FROM alpine:latest

WORKDIR /app/

ARG LOG_DIR=/app/logs

RUN mkdir -p ${LOG_DIR}

ENV LOGGER_FILE=${LOG_DIR}/finance.log

COPY --from=builder /app/bin/bot .
COPY --from=builder /app/configs configs/
COPY --from=builder /app/migrations migrations/

VOLUME ${LOG_DIR}

ENTRYPOINT ["./bot"]