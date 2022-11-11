package main

import (
	"os"

	"github.com/open-policy-agent/opa/cmd"
	"github.com/open-policy-agent/opa/runtime"
	"github.com/tibotix/opa-redis-plugin/plugin"
)

func main() {
	runtime.RegisterPlugin(plugin.PluginName, plugin.Factory{})

	if err := cmd.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
