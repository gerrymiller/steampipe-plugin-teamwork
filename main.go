package main

import (
	"steampipe-plugin-teamwork/teamwork"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: teamwork.Plugin})
}
