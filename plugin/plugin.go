package plugin

import (
	"github.com/open-policy-agent/opa/plugins"

	"github.com/tibotix/opa-redis-plugin/internal"
)

// Factory defines the interface OPA uses to instantiate a plugin.
type Factory struct{}

// PluginName is the name to register with the OPA plugin manager
const PluginName = internal.PluginName

// New returns the object initialized with a valid plugin configuration.
func (Factory) New(m *plugins.Manager, config any) plugins.Plugin {
	return internal.New(m, config.(*internal.ParsedConfig))
}

// Validate returns a valid configuration to instantiate the plugin.
func (Factory) Validate(m *plugins.Manager, config []byte) (interface{}, error) {
	return internal.Validate(m, config)
}
