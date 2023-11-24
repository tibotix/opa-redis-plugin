package plugin

import (
	"github.com/open-policy-agent/opa/plugins"

	"github.com/tibotix/opa-redis-plugin/internal"
	"github.com/tibotix/opa-redis-plugin/redisManager"
)

// Factory defines the interface OPA uses to instantiate a plugin.
type Factory struct {
	redisManager *redisManager.RedisManager
}

func NewFactory(rm *redisManager.RedisManager) Factory {
	return Factory{
		rm,
	}
}

// PluginName is the name to register with the OPA plugin manager
const PluginName = internal.PluginName

// New returns the object initialized with a valid plugin configuration.
func (f Factory) New(m *plugins.Manager, config any) plugins.Plugin {
	m.UpdatePluginStatus(PluginName, &plugins.Status{State: plugins.StateNotReady})
	return internal.New(m, f.redisManager, config.(*internal.ParsedConfig))
}

// Validate returns a valid configuration to instantiate the plugin.
func (f Factory) Validate(m *plugins.Manager, config []byte) (interface{}, error) {
	return internal.Validate(m, config)
}
