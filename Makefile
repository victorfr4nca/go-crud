run-api:
	air -c cmd/api/.air.toml

build:
	go build -o bin/main cmd/main.go

test-verbose:
	go test -v ./...

test:
	go test ./...
