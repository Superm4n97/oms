push-auth-server:
	@docker build -f ./auth-server/Dockerfile -t superm4n/auth-server .
	@docker push superm4n/auth-server

push-oms-server:
	@docker build -f ./oms/Dockerfile -t superm4n/order-server .
	@docker push superm4n/order-server

push-order-processor:
	@docker build -f ./order-processor/Dockerfile -t superm4n/order-processor .
	@docker push superm4n/order-processor

push-all: push-auth-server push-oms-server push-order-processor

ADDTL_LINTERS   := goconst,gofmt,goimports,unparam

OS   := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

GO_VERSION       ?= 1.20
BUILD_IMAGE      ?= ghcr.io/appscode/golang-dev:$(GO_VERSION)

.PHONY: lint
lint: $(BUILD_DIRS)
	@echo "running linter"
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    --env GO111MODULE=on                                    \
	    --env GOFLAGS="-mod=vendor"                             \
	    $(BUILD_IMAGE)                                          \
	    golangci-lint run --enable $(ADDTL_LINTERS) --timeout=10m --skip-files="generated.*\.go$\" --skip-dirs-use-default