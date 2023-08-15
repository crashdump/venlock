.PHONY: build clean test
.DEFAULT_GOAL=test

build: test
	go build -v -o dist/libguardian cmd/libguardian/*.go

clean:
	rm -f dist/libguardian

test:
	go mod tidy
	staticcheck ./...
	go test -v ./...

integration-test:
	go test -v ./... -tags=integration

release: test integration-test build
	go test all -v