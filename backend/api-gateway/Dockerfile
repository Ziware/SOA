# Этап сборки
FROM golang:1.24.1-alpine

WORKDIR /app
COPY . .

WORKDIR /app/api-gateway/server
RUN go mod download -x
RUN go build

ENTRYPOINT ["./server"]
