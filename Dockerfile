FROM golang:1.24-bookworm AS builder

ENV GOFLAGS="-mod=readonly"

RUN apt-get update && apt-get install -y --no-install-recommends \
    bash make git curl build-essential pkg-config libssl-dev ca-certificates \
    && curl https://sh.rustup.rs -sSf | bash -s -- -y \
    && . "$HOME/.cargo/env"
ENV PATH="/root/.cargo/bin:${PATH}"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM debian:bookworm-slim

COPY --from=builder /app/opa_redis_plugin /opa_redis_plugin
ENTRYPOINT [ "/opa_redis_plugin" ]
CMD ["run"]
