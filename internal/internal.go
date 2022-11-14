package internal

import (
	"context"
	"errors"
	"time"

	"github.com/open-policy-agent/opa/plugins"
	"github.com/open-policy-agent/opa/util"

	"github.com/tibotix/opa-redis-plugin/proxy"

	"github.com/go-redis/redis/v8"
)

const (
	PluginName         = "redis"
	defaultEnabled     = true
	defaultAddress     = ""
	defaultMaxRetries  = 3
	defaultDialTimeout = 8
	defaultReadTimeout = 2
)

type config struct {
	Enabled     bool          `json:"enabled"`
	Address     string        `json:"address"`
	MaxRetries  int           `json:"max_retries"`
	DialTimeout time.Duration `json:"dial_timeout_in_seconds"`
	ReadTimeout time.Duration `json:"read_timeout_in_seconds"`
	// TODO: support TLS config
}

type ParsedConfig struct {
	Enabled bool
	Options *redis.Options
}

type redisPlugin struct {
	parsedConfig ParsedConfig
	redisProxy   proxy.Proxy[*redis.Client]
	redisContext context.Context
}

func Validate(m *plugins.Manager, bs []byte) (*ParsedConfig, error) {
	config := config{
		Enabled:     defaultEnabled,
		Address:     defaultAddress,
		MaxRetries:  defaultMaxRetries,
		DialTimeout: defaultDialTimeout,
		ReadTimeout: defaultReadTimeout,
	}

	if err := util.Unmarshal(bs, &config); err != nil {
		return nil, err
	}

	if config.Address == defaultAddress {
		return nil, errors.New("No Redis Address Provided.")
	}

	opt, err := redis.ParseURL(config.Address)
	if err != nil {
		return nil, err
	}
	opt.MaxRetries = config.MaxRetries
	opt.DialTimeout = time.Second * config.DialTimeout
	opt.ReadTimeout = time.Second * config.ReadTimeout

	parsedConfig := ParsedConfig{
		Enabled: config.Enabled,
		Options: opt,
	}

	return &parsedConfig, nil

}

func New(m *plugins.Manager, parsedConfig *ParsedConfig) plugins.Plugin {
	plugin := &redisPlugin{
		parsedConfig: *parsedConfig,
		redisContext: context.Background(),
	}

	return plugin
}

func (p *redisPlugin) startRedisClient() {
	p.redisProxy.Set(redis.NewClient(p.parsedConfig.Options))
}

func (p *redisPlugin) Start(ctx context.Context) error {
	if !p.parsedConfig.Enabled {
		return nil
	}
	go p.startRedisClient()

	p.registerManualCommands()
	p.registerAutogenCommands()

	return nil
}

func (p *redisPlugin) Stop(ctx context.Context) {
	rdb, err := p.redisProxy.Get()
	if err != nil {
		return
	}
	rdb.Close()
	p.redisProxy.Unset()
}

func (p *redisPlugin) Reconfigure(ctx context.Context, parsedConfig interface{}) {
	// TODO: add testing using redismock and test reconfigure
	p.Stop(ctx)
	p.parsedConfig = parsedConfig.(ParsedConfig)
	go p.startRedisClient()
}
