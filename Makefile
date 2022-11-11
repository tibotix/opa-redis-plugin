

BIN := opa_redis

CGO_ENABLED ?= 1
WASM_ENABLED ?= 1

GO := CGO_ENABLED=$(CGO_ENABLED) GOFLAGS=-mod=vendor GOPROXY=off go

GO_TAGS := -tags=
ifeq ($(WASM_ENABLED),1)
GO_TAGS = -tags=opa_wasm
endif

build:
	$(GO) build $(GO_TAGS) -o $(BIN) ./cmd/opa-redis-plugin/...
# TODO: Add ldflags 