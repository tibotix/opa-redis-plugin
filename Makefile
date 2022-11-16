VERSION_OPA := $(shell ./scripts/get-opa-version.sh)
VERSION := $(VERSION_OPA)-redis$(shell ./scripts/get-plugin-rev.sh)

CGO_ENABLED ?= 1
WASM_ENABLED ?= 1

GO := CGO_ENABLED=$(CGO_ENABLED) GOFLAGS=-mod=vendor GOPROXY=off go
GOVERSION := $(shell ./scripts/get-go-version.sh)
GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
DISABLE_CGO := CGO_ENABLED=0

BIN := opa_redis_$(GOOS)_$(GOARCH)

GO_TAGS := -tags=
ifeq ($(WASM_ENABLED),1)
GO_TAGS = -tags=opa_wasm
endif


BUILD_COMMIT := $(shell ./scripts/get-build-commit.sh)
BUILD_TIMESTAMP := $(shell ./scripts/get-build-timestamp.sh)
BUILD_HOSTNAME := $(shell ./scripts/get-build-hostname.sh)

LDFLAGS := "-X github.com/open-policy-agent/opa/version.Version=$(VERSION) \
	-X github.com/open-policy-agent/opa/version.Vcs=$(BUILD_COMMIT) \
	-X github.com/open-policy-agent/opa/version.Timestamp=$(BUILD_TIMESTAMP) \
	-X github.com/open-policy-agent/opa/version.Hostname=$(BUILD_HOSTNAME)"


generate:
	./scripts/gen-redis-commands.py impl > ./internal/redis_commands_autogen.go
	go fmt ./internal/redis_commands_autogen.go

generate-doc:
	./scripts/gen-redis-commands.py doc > ./doc/supported_commands.md

build:
	$(GO) build $(GO_TAGS) -o $(BIN) -ldflags $(LDFLAGS) ./cmd/opa-redis-plugin/...
