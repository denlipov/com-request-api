.PHONY: build
build:
	go build cmd/com-request-api/main.go

.PHONY: test
test:
	go test -count=1 -v ./...
