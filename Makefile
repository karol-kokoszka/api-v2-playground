export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)
export OS := $(shell uname)

all: help

.PHONY: gen
gen: ## Generate
	@sudo rm -Rf gen
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/api/publicapi.openapi.json -g go-gin-server -o /local/gen/external --additional-properties=packageName=restapi,apiPath=restapi
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/api/publicapi.openapi.json -g go -o /local/gen/sdk --additional-properties=generateInterfaces=true

GO = CGO_ENABLED=0 go build -mod=mod

.PHONY: build
build: ## Build project binary
	@$(GO) -o ./apiv2-playground .

gen-server:
	oapi-codegen -generate chi-server,types -package siren api/publicapi.openapi.json > api/gen/siren/api.gen.go        