ARG GO_VERSION=""
FROM golang:1.12.0-alpine3.9
# FROM golang:${GO_VERSION:+${GO_VERSION}-}alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main .

CMD ["/app/main"]
