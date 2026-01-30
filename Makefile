BIN="./bin"

.PHONY: test lint build generate-buf

test:
	$(info ******************** running tests ********************)
	@go test -v ./...

lint:
	$(info ******************** running lint tools ********************)
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run -v

build: gen-buf

gen-buf:
	$(info ******************** generating buf ********************)
	go run github.com/bufbuild/buf/cmd/buf generate

lint-buf:
	$(info ******************** running lint tool for buf ********************)
	go run github.com/bufbuild/buf/cmd/buf lint
