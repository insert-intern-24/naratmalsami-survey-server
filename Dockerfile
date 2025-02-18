FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o myapp .

FROM alpine:latest

COPY --from=builder /app/myapp /myapp

CMD ["/myapp"]
