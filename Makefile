# Fetch OS info
GOVERSION=$(shell go version)
UNAME_OS=$(shell go env GOOS)
UNAME_ARCH=$(shell go env GOARCH)

BUF_VERSION:= 1.5.0

BUILD_OUT_DIR := "bin/"

API_OUT       := "bin/api"
API_MAIN_FILE := "cmd/api/main.go"

MOCK_IN := "internal/loans/mock"

RPC_ROOT := "rpc/"

.PHONY: go-build-api ## Build the binary file for API server
go-build-api:
	@echo "\n Building Loans API"
	@CGO_ENABLED=0 GOOS=$(UNAME_OS) GOARCH=$(UNAME_ARCH) go build -v -o $(API_OUT) $(API_MAIN_FILE)

.PHONY: test-unit
test-unit:
	go test -tags=unit -timeout 5m `go list ./...` -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -func=coverage.out

.PHONY: proto-deps
proto-deps:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	@go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.0
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

	curl -sSL \
	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
	-o "$(shell go env GOPATH)/bin/buf" && \
	chmod +x "$(shell go env GOPATH)/bin/buf"

	buf mod update

.PHONY: proto-refresh ## Download and re-compile protobuf
proto-refresh: clean proto-generate ## Fetch proto files frrm remote repo

.PHONY: proto-generate ## Compile protobuf to pb files
proto-generate:
	@echo "\n + Generating pb language bindings\n"
	buf generate

.PHONY: clean ## Remove previous builds, protobuf files, and proto compiled code
clean:
	 @echo " + Removing cloned and generated files\n"
	 @rm -rf $(API_OUT) $(RPC_ROOT)

.PHONY: mock-deps
mock-deps:
	@echo "\n + Fetching mocking related dependencies \n"
	@go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: generate-mocks ## Generate mocks from interfaces
generate-mocks:
	@echo "\n Clearing any existing mocks in $(MOCK_IN)\n"
	rm -rf $(MOCK_IN)
	rm -rf $(PKG_MOCK_IN)
	@echo "\n Generating new mocks in $(MOCK_IN)\n"
	mockgen -destination=internal/loan/mock/repository.go -package=mock github.com/gaganchawara/loans/internal/loan/interfaces Repository
