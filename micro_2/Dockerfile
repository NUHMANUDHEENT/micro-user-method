FROM golang:1.22.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .

RUN go build -o micro-2 ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/micro-2 .

COPY .env ./

CMD ["./micro-2"]