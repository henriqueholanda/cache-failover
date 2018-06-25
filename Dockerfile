FROM golang:1.10-alpine

RUN apk add --update git

WORKDIR /go/src/github.com/henriqueholanda/cache-failover

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN go build -o cache-failover

EXPOSE 80

CMD ["./cache-failover"]
