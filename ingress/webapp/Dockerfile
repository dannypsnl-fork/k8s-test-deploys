FROM golang:1.10

WORKDIR /go/src/webapp
COPY . .
RUN go get -t ./...
RUN go build
WORKDIR /go/bin

ENTRYPOINT ./webapp
