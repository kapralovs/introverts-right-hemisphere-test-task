# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o simple-test-api ./cmd/main.go

EXPOSE 8080

CMD [ "./simple-test-api" ]