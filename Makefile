VERSION_OPA := $(shell ./build/get-opa-version.sh)
VERSION := $(VERSION_OPA)-redis$(shell ./build/get-plugin-rev.sh)

BIN := opa_redis

CGO_ENABLED ?= 1
WASM_ENABLED ?= 1

GO := CGO_ENABLED=$(CGO_ENABLED) GOFLAGS=-mod=vendor GOPROXY=off go

GO_TAGS := -tags=
ifeq ($(WASM_ENABLED),1)
GO_TAGS = -tags=opa_wasm
endif


BUILD_COMMIT := $(shell ./build/get-build-commit.sh)
BUILD_TIMESTAMP := $(shell ./build/get-build-timestamp.sh)
BUILD_HOSTNAME := $(shell ./build/get-build-hostname.sh)

LDFLAGS := "-X github.com/open-policy-agent/opa/version.Version=$(VERSION) \
	-X github.com/open-policy-agent/opa/version.Vcs=$(BUILD_COMMIT) \
	-X github.com/open-policy-agent/opa/version.Timestamp=$(BUILD_TIMESTAMP) \
	-X github.com/open-policy-agent/opa/version.Hostname=$(BUILD_HOSTNAME)"

build:
	$(GO) build $(GO_TAGS) -o $(BIN) -ldflags $(LDFLAGS) ./cmd/opa-redis-plugin/...