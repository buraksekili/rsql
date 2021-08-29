testcover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

build:
	go build -o $(GOPATH)/bin/rsql

test:
	go test ./... -v -cover