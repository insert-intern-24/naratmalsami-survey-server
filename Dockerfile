FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY .env .env
COPY . .

RUN go mod tidy && go build -o myapp .

FROM alpine:latest

COPY --from=builder /app/myapp /myapp
COPY --from=builder /app/.env /app/.env

CMD ["/myapp"]