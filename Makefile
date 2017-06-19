.DEFAULT_GOAL := help

PROJECT_NAME := NEO4BACON

.PHONY: help prep image deploy test run

help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

grpc: ## generate gRPC code
	@protoc -I . neo4bacon.proto --go_out=plugins=grpc:api/

prep: ## download dependencies and format code
	@go get -d ./...
	@go fmt ./...

image: ## build docker image and push image to registry
	@docker build -t alesr/neo4bacon -f resources/prod/Dockerfile .
	@docker tag alesr/neo4bacon alesr/neo4bacon:latest
	@docker push alesr/neo4bacon

run: ## deploy application container
	@docker run --rm -d -p 50051:50051 --name neo4bacon alesr/neo4bacon

build/osx: prep ## build osx binary
	@go build

build/linux: prep ## build linux binary
	@GOOS=linux go build

test: ## run unit tests
	@docker build -t neo4bacon-test -f resources/test/Dockerfile .
	@docker run --rm neo4bacon-test
