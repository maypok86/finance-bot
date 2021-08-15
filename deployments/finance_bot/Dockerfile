FROM golang:1.16-alpine AS builder
LABEL stage=builder

RUN go version

COPY . /github.com/LazyBearCT/finance-bot/
WORKDIR /github.com/LazyBearCT/finance-bot/

RUN go mod download
RUN GOOS=linux go build -o ./bin/bot ./cmd/bot

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /github.com/LazyBearCT/finance-bot/bin/bot .
COPY --from=builder /github.com/LazyBearCT/finance-bot/configs configs/
COPY --from=builder /github.com/LazyBearCT/finance-bot/migrations migrations/

ENTRYPOINT ["./bot"]