run-api:
	air -c cmd/api/.air.toml

build:
	go build -o bin/main cmd/main.go

test:
	go test -v ./...
