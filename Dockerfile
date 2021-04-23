FROM golang:buster

COPY . /app
WORKDIR /app

RUN go build cmd/http/main.go

ENTRYPOINT ["./main"]
