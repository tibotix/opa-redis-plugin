package main

import (
	"os"

	"github.com/open-policy-agent/opa/cmd"
	"github.com/open-policy-agent/opa/runtime"
	"github.com/tibotix/opa-redis-plugin/plugin"
	"github.com/tibotix/opa-redis-plugin/redisManager"
)

func main() {
	redisManager := redisManager.New()
	redisManager.RegisterCommands()

	runtime.RegisterPlugin(plugin.PluginName, plugin.NewFactory(redisManager))

	if err := cmd.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
