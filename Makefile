BIN="./bin"

.PHONY: test lint

test:
	$(info ******************** running tests ********************)
	@go test -v ./...

lint:
	$(info ******************** running lint tools ********************)
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run -v
