FROM golang:1.10-alpine

RUN apk add --update git

WORKDIR /go/src/github.com/henriqueholanda/cache-failover

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

ENV API_PORT=80

RUN go build -o cache-failover

EXPOSE 80

CMD ["./cache-failover"]
