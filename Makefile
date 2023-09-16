BUILD_OUT_DIR := "bin/"

API_OUT       := "bin/api"
API_MAIN_FILE := "cmd/api/main.go"
# Fetch OS info
GOVERSION=$(shell go version)
UNAME_OS=$(shell go env GOOS)
UNAME_ARCH=$(shell go env GOARCH)

.PHONY: go-build-api ## Build the binary file for API server
go-build-api:
	@echo "\n Building Loans API"
	@CGO_ENABLED=0 GOOS=$(UNAME_OS) GOARCH=$(UNAME_ARCH) go build -v -o $(API_OUT) $(API_MAIN_FILE)

.PHONY: test-unit
test-unit:
	go test -tags=unit -timeout 5m `go list ./...` -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -func=coverage.out
