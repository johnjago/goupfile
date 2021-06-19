all: build run

ci: build test

build:
	go build .

test:
	go test

run:
	./server

clean:
	rm -rf uploads/ server sqlite_db
