GO_FLAGS   ?=
BUILD_PATH ?= "./"
NAME       := name-ddns
OUTPUT_BIN ?= ${BUILD_PATH}${NAME}
VERSION    ?= v0.1.0
IMG_NAME   := naxhh/name-ddns
IMAGE      := ${IMG_NAME}:${VERSION}
IMAGE_M1   := ${IMAGE}-arm64

default: help

fmt: ## Formats all files
	@go fmt ./...

test: ## Run all tests
	@go clean --testcache && go test ./...

cover: ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

build: ## Builds the CLI
	@go build ${GO_FLAGS} \
	-o ${OUTPUT_BIN} \
	./cmd/ddns

docker: ## Build Docker Image
	@docker build \
	--platform=linux/amd64 \
	--rm -t ${IMAGE} .

	@docker build \
	--platform=linux/arm64 \
	--rm -t ${IMAGE_M1} .

publish: ## Publishes docker images
	@docker tag ${IMAGE} ${IMG_NAME}:latest
	@docker push ${IMG_NAME}:latest
	@docker push ${IMAGE}
	@docker push ${IMAGE_M1}

help: ## Displays this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'