all: build run

ci: build test

build:
	go build -v ./...

test:
	go test -v ./...

run:
	./server

clean:
	rm -rf uploads/ server sqlite_db
