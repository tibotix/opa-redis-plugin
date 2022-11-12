Table of Contents
- [Overview](#overview)
- [Usage](#usage)
- [Configuration](#configuration)
- [Example](#example)
- [Dependencies](#dependencies)
- [Currently Implemented Commands](#currently-implemented-commands)
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

    opa_redis_plugin "github.com/tibotix/opa-redis-plugin"
)

func main() {
	runtime.RegisterPlugin(opa_redis_plugin.PluginName, opa_redis_plugin.Factory{})

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

- CLIENTGETNAME
- ECHO
- PING
- QUIT
- DEL
- UNLINK
- DUMP
- EXISTS
- EXPIRE
- EXPIREAT
- EXPIRENX
- EXPIREXX
- EXPIREGT
- EXPIRELT
- KEYS
- MIGRATE
- MOVE
- OBJECTREFCOUNT
- OBJECTENCODING
- OBJECTIDLETIME
- PERSIST
- PEXPIRE
- PEXPIREAT
- PTTL
- RANDOMKEY
- RENAME
- RENAMENX
- RESTORE
- RESTOREREPLACE
- TOUCH
- TTL
- TYPE
- APPEND
- DECR
- DECRBY
- GET
- GETRANGE
- GETSET
- GETEX
- GETDEL
- INCR
- INCRBY
- INCRBYFLOAT
- MSET
- MSETNX
- SET
- SETEX
- SETNX
- SETXX
- SETRANGE
- STRLEN
- COPY
- GETBIT
- SETBIT
- BITOPAND
- BITOPOR
- BITOPXOR
- BITOPNOT
- BITPOS
- BITFIELD
- HDEL
- HEXISTS
- HGET
- HINCRBY
- HINCRBYFLOAT
- HKEYS
- HLEN
- HSET
- HMSET
- HSETNX
- HVALS
- HRANDFIELD
- BLPOP
- BRPOP
- BRPOPLPUSH
- LINDEX
- LINSERT
- LINSERTBEFORE
- LINSERTAFTER
- LLEN
- LPOP
- LPOPCOUNT
- LPUSH
- LPUSHX
- LRANGE
- LREM
- LSET
- LTRIM
- RPOP
- RPOPCOUNT
- RPOPLPUSH
- RPUSH
- RPUSHX
- LMOVE
- BLMOVE
- SADD
- SCARD
- SDIFF
- SDIFFSTORE
- SINTER
- SINTERSTORE
- SISMEMBER
- SMISMEMBER
- SMEMBERS
- SMOVE
- SPOP
- SPOPN
- SRANDMEMBER
- SRANDMEMBERN
- SREM
- SUNION
- SUNIONSTORE
- XDEL
- XLEN
- XGROUPCREATE
- XGROUPCREATEMKSTREAM
- XGROUPSETID
- XGROUPDESTROY
- XGROUPCREATECONSUMER
- XGROUPDELCONSUMER
- XACK
- XTRIM
- XTRIMAPPROX
- XTRIMMAXLEN
- XTRIMMAXLENAPPROX
- XTRIMMINID
- XTRIMMINIDAPPROX
- ZADD
- ZADDNX
- ZADDXX
- ZADDCH
- ZADDNXCH
- ZADDXXCH
- ZCARD
- ZCOUNT
- ZLEXCOUNT
- ZINCRBY
- ZMSCORE
- ZRANGE
- ZRANK
- ZREM
- ZREMRANGEBYRANK
- ZREMRANGEBYSCORE
- ZREMRANGEBYLEX
- ZREVRANGE
- ZREVRANK
- ZSCORE
- ZRANDMEMBER
- ZDIFF
- ZDIFFSTORE
- PFADD
- PFCOUNT
- PFMERGE
- BGREWRITEAOF
- BGSAVE
- CLIENTKILL
- CLIENTKILLBYFILTER
- CLIENTLIST
- CLIENTPAUSE
- CLIENTID
- CONFIGRESETSTAT
- CONFIGSET
- CONFIGREWRITE
- DBSIZE
- FLUSHALL
- FLUSHALLASYNC
- FLUSHDB
- FLUSHDBASYNC
- INFO
- LASTSAVE
- SAVE
- SHUTDOWN
- SHUTDOWNSAVE
- SHUTDOWNNOSAVE
- SLAVEOF
- TIME
- DEBUGOBJECT
- READONLY
- READWRITE
- MEMORYUSAGE
- SCRIPTEXISTS
- SCRIPTFLUSH
- SCRIPTKILL
- SCRIPTLOAD
- PUBLISH
- PUBSUBCHANNELS
- PUBSUBNUMPAT
- CLUSTERNODES
- CLUSTERMEET
- CLUSTERFORGET
- CLUSTERREPLICATE
- CLUSTERRESETSOFT
- CLUSTERRESETHARD
- CLUSTERINFO
- CLUSTERKEYSLOT
- CLUSTERGETKEYSINSLOT
- CLUSTERCOUNTFAILUREREPORTS
- CLUSTERCOUNTKEYSINSLOT
- CLUSTERDELSLOTS
- CLUSTERDELSLOTSRANGE
- CLUSTERSAVECONFIG
- CLUSTERSLAVES
- CLUSTERFAILOVER
- CLUSTERADDSLOTS
- CLUSTERADDSLOTSRANGE
- GEOADD
- GEODIST
- GEOHASH
  
More Information about all commands can be found [here](https://redis.io/commands)


# Credits

Special thanks to the whole OPA team for maintaining OPA and making it open-source,
and the [opa-envoy-plugin](https://github.com/open-policy-agent/ope-envoy-plugin) maintainer
for a good plugin template and code inspirations that this plugin is heavily based on.
