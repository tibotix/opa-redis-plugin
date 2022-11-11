package internal

import (
	"context"
	"time"

	"github.com/open-policy-agent/opa/plugins"
	"github.com/open-policy-agent/opa/util"

	"github.com/tibotix/opa-redis-plugin/proxy"

	"github.com/go-redis/redis/v8"
)

const (
	PluginName         = "redis"
	defaultEnabled     = true
	defaultAddress     = "redis://localhost:6379/0"
	defaultMaxRetries  = 3
	defaultDialTimeout = 8 * time.Second
	defaultReadTimeout = 2 * time.Second
)

type Config struct {
	Enabled     bool          `json:"enabled"`
	Address     string        `json:"address"`
	MaxRetries  int           `json:"max_retries"`
	DialTimeout time.Duration `json:"dial_timeout"`
	ReadTimeout time.Duration `json:"read_timeout"`
	// TODO: support TLS config
}

type redisPlugin struct {
	cfg          Config
	redisProxy   proxy.Proxy[*redis.Client]
	redisContext context.Context
}

func Validate(m *plugins.Manager, bs []byte) (*Config, error) {
	cfg := Config{
		Enabled:     defaultEnabled,
		Address:     defaultAddress,
		MaxRetries:  defaultMaxRetries,
		DialTimeout: defaultDialTimeout,
		ReadTimeout: defaultReadTimeout,
	}

	if err := util.Unmarshal(bs, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}

func New(m *plugins.Manager, cfg *Config) plugins.Plugin {
	plugin := &redisPlugin{
		cfg:          *cfg,
		redisContext: context.Background(),
	}

	return plugin
}

func (p *redisPlugin) startRedisClient() {
	opt, err := redis.ParseURL(p.cfg.Address)
	if err != nil {
		return
	}
	opt.MaxRetries = p.cfg.MaxRetries
	opt.DialTimeout = p.cfg.DialTimeout
	p.redisProxy.Set(redis.NewClient(opt))
}

func (p *redisPlugin) Start(ctx context.Context) error {
	if !p.cfg.Enabled {
		return nil
	}
	go p.startRedisClient()

	p.registerGET()
	p.registerEXISTS()
	p.registerSISMEMBER()

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

func (p *redisPlugin) Reconfigure(ctx context.Context, cfg interface{}) {
	p.Stop(ctx)
	p.cfg = cfg.(Config)
	go p.startRedisClient()
}
