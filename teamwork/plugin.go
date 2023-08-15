package teamwork

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const pluginName = "steampipe-plugin-teamwork"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		//DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"teamwork_project": tableTeamworkProject(ctx),
		},
	}
	return p
}
