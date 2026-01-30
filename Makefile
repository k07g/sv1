BIN="./bin"

.PHONY: test lint build generate-buf

test:
	$(info ******************** running tests ********************)
	@go test -v ./...

lint: lint-go lint-buf

build: gen-buf

lint-go:
	$(info ******************** running lint tools for Go ********************)
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run

gen-buf:
	$(info ******************** generating buf ********************)
	go run github.com/bufbuild/buf/cmd/buf generate

lint-buf:
	$(info ******************** running lint tool for buf ********************)
	@go run github.com/bufbuild/buf/cmd/buf lint
