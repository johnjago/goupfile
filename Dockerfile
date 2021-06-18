FROM golang:1.16-buster

WORKDIR /go/src/server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["server"]
