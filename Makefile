run:
	@go run .

lint:
	@golangci-lint run

test:
	go test ./... -v -count 1
