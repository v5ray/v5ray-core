package plugins

import "github.com/v4fly/v4ray-core/v0/main/commands/base"

var Plugins []Plugin

type Plugin func(*base.Command) func() error

func RegisterPlugin(plugin Plugin) {
	Plugins = append(Plugins, plugin)
}
