BIN="./bin"

.PHONY: test lint build fmt lint-go gen-buf lint-buf

test:
	$(info ******************** running tests ********************)
	@go test -v ./...

lint: lint-go lint-buf

build: gen-buf fmt
	$(info ******************** building binary ********************)
	@go build -o ./bin/sv1 ./cmd/app

fmt:
	$(info ******************** formatting files ********************)
	@go fmt ./...
	@go run github.com/bufbuild/buf/cmd/buf format -w
	@go run golang.org/x/tools/cmd/goimports -w ./
	@go run mvdan.cc/gofumpt -l -w .

lint-go:
	$(info ******************** running lint tools for Go ********************)
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run

gen-buf:
	$(info ******************** generating buf ********************)
	go run github.com/bufbuild/buf/cmd/buf generate

lint-buf:
	$(info ******************** running lint tool for buf ********************)
	@go run github.com/bufbuild/buf/cmd/buf lint
