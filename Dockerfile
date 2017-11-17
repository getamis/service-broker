# Build broker in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . $GOPATH/src/github.com/getamis/service-broker
RUN cd $GOPATH/src/github.com/getamis/service-broker && make broker && mv build/bin/broker /broker

# Pull broker into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /broker /usr/local/bin/

ENTRYPOINT ["broker"]