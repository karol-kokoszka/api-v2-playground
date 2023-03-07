export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)
export OS := $(shell uname)

all: help

.PHONY: install-dependencies
install-dependencies:
	@rm -Rf bin
	cd ./dev-tools; go install ./...
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/go-bridget/twirp-swagger-gen/cmd/protoc-gen-twirp-swagger@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest
ifeq ($(OS),Darwin)
	brew install skaffold helm
endif

.PHONY: gen
gen: ## Generate
	@sudo rm -Rf gen
	buf mod update
	buf generate --path=proto/srv
	buf generate --path=proto/srv --template=buf.gen.mock.yaml
	buf generate --path=proto/restapi --template=buf.gen.openapiv2.yaml
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/api/publicapi.openapi.json -g go-gin-server -o /local/gen/external --additional-properties=packageName=restapi,apiPath=restapi
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/api/publicapi.openapi.json -g go -o /local/gen/sdk --additional-properties=generateInterfaces=true

.PHONY: fmt
fmt: ## Format source code
	@$(GOBIN)/golangci-lint run -c .golangci-fmt.yml --fix

.PHONY: check
check:
	@buf lint --path=proto/srv
	@buf lint --path=proto/restapi
	@$(GOBIN)/golangci-lint run
	@go vet ./...

.PHONY: breaking
breaking:
	@buf breaking --against '.git#branch=main'

.PHONY: vendor
vendor: ## Fix dependencies and make vendored copies
	@go mod tidy -compat=1.17
	@go mod vendor

GO = CGO_ENABLED=0 go build -mod=mod

.PHONY: build
build: ## Build project binary
	@$(GO) -o ./scylla-cloud .

.PHONY: k8s-build
k8s-build: ## Build/link binary for k8s
ifeq ($(OS),Linux)
	@$(GO) -o ./scylla-cloud .
	@ln -sf ./scylla-cloud ./scylla-cloud-linux
else ifeq ($(OS),Darwin)
	@GOOS=linux $(GO) -o ./scylla-cloud-linux .
else
	$(error Unsupported OS $(OS))
endif

.PHONY: start-dev-env
start-dev-env: k8s-build
	skaffold config set --global local-cluster true
	@skaffold dev

.PHONY: test
test: ## Run unit tests
	@go test -tags=mock ./...

.PHONY: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {printf "\033[36m%-25s\033[0m %s\n", $$1, $$NF}' Makefile

install-oapi:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

gen-server:
	oapi-codegen -generate chi-server,types -package siren api/publicapi.openapi.json > api/gen/siren/api.gen.go        