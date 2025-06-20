FROM golang:1.22-alpine

WORKDIR /app

COPY cmd/server/main.go .

RUN go build -o main main.go

CMD ["./main"]
