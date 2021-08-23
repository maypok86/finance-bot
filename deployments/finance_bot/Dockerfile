FROM golang:1.16-alpine AS builder
LABEL stage=builder

RUN go version

COPY . /gitlab.com/LazyBearCT/finance-bot/
WORKDIR /gitlab.com/LazyBearCT/finance-bot/

RUN go mod download
RUN GOOS=linux go build -o ./bin/bot ./cmd/bot

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /gitlab.com/LazyBearCT/finance-bot/bin/bot .
COPY --from=builder /gitlab.com/LazyBearCT/finance-bot/configs configs/
COPY --from=builder /gitlab.com/LazyBearCT/finance-bot/migrations migrations/

ENTRYPOINT ["./bot"]