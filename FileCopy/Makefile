all: linters tests build

build: FileCopy.go
	go build FileCopy.go
tests: FileCopy_test.go
	go test -v
linters:
	golangci-lint run
