BIN="./bin"

.PHONY: test

test:
	$(info ******************** running tests ********************)
	go test -v ./...