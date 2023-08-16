package teamwork

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type teamworkConfig struct {
	APIKey *string `cty:"api_key"`
	Domain *string `cty:"domain"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"domain": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &teamworkConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) teamworkConfig {
	if connection == nil || connection.Config == nil {
		return teamworkConfig{}
	}
	config, _ := connection.Config.(teamworkConfig)
	return config
}
