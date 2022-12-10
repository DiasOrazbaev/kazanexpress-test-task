lint:
	gofmt -w -s ./..
	golangci-lint run --fix

run:
	go run cmd/service/main.go &
	go run cmd/client/main.go

.PHONY: lint run