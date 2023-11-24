
![opa-redis-plugin](doc/opa-redis-plugin.png)


- [Overview](#overview)
- [Usage](#usage)
- [Configuration](#configuration)
- [Example](#example)
- [Dependencies](#dependencies)
- [Currently Implemented Commands](#currently-implemented-commands)
- [Executing Arbitrary Commands](#executing-arbitrary-commands)
- [Limitations](#limitations)
- [Credits](#credits)

# Overview

This repository contains an [OPA](https://github.com/open-policy-agent/opa) plugin that
provides [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) bindings for the [go-redis](https://github.com/go-redis/redis) library.


# Usage

You can use this plugin as a library to embed in your go code:
```golang
import (
    "os"

    "github.com/open-policy-agent/opa/cmd"
    "github.com/open-policy-agent/opa/runtime"

    opa_redis_plugin "github.com/tibotix/opa-redis-plugin/plugin"
)

func main() {
	redisManager := opa_redis_plugin.NewRedisManager()
	redisManager.RegisterCommands()

	runtime.RegisterPlugin(opa_redis_plugin.PluginName, opa_redis_plugin.NewFactory(redisManager))

	if err := cmd.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
```

# Configuration

This plugin expects some configuration in the configuration file provided to opa.
For example:

`opa run -c config.yaml`

config.yaml:
```yaml
plugins:
  redis:
    enabled: true # DEFAULT: true
    address: "redis://${REDIS_USERNAME}:${REDIS_PASSWORD}@${REDIS_HOST}:${REDIS_PORT}/0" # REQUIRED
    max_retries: 3 # DEFAULT: 3
    dial_timeout_in_seconds: 8.0 # DEFAULT: 8
    read_timeout_in_seconds: 2.0 # DEFAULT: 2
```

# Example

When this plugin is enabled, you can use all `redis.*` functions.
Each function name is the lowercase version of the corresponding command, e.g.

    CLIENTGETNAME -> redis.clientgetname
    GET -> redis.get
    SMISMBER -> redis.smismember
    [...]

example.rego:
```py
redis.get("key")
redis.smismember("set_key", ["item1", "item2"])
```
NOTE: On Connection Error all redis calls return undefined.

# Dependencies

This plugin depends on the following modules:
- `"github.com/go-redis/redis v8.11.5"`
- `"github.com/open-policy-agent/opa v0.46.1"`

# Currently Implemented Commands

Please see [supported_commands.md](./doc/supported_commands.md)

# Executing Arbitrary Commands

To execute arbitrary commands, use the `redis.do` function:

```
redis.do(["SET", "mykey", "myvalue"])
```

# Limitations

Currently, Redis [Transactions](https://redis.io/docs/manual/transactions/) and [Pipelines](https://redis.io/docs/manual/pipelining/) are not supported yet. Pull Requests are welcome.

Additionally, `opa eval` does not work with this plugin, as `opa eval` does not load plugins.

# Credits

Special thanks to the whole OPA team for maintaining OPA and making it open-source,
 the [opa-envoy-plugin](https://github.com/open-policy-agent/ope-envoy-plugin) maintainer
for a good plugin template and code inspirations that this plugin is based on and the
[go-redis](https://github.com/go-redis/redis) library, which is used internally.
