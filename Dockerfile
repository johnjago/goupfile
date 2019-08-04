FROM golang:buster

WORKDIR /go/src/core
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["core"]
