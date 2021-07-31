all: run

ci: build test

build:
	go build -v ./...

test:
	go test -v ./...

run:
	go run -v ./...

clean:
	rm -rf uploads/ server sqlite_db
