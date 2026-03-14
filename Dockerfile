FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod downland

COPY . .
RUN go build -o bot .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bot .

CMD [ "./bot" ]