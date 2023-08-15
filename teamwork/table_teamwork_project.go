package teamwork

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTeamworkProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "teamwork_project",
		Description: "Projects from Teamwork.com",
		List: &plugin.ListConfig{
			Hydrate: listTeamworkProjects,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the project.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the project.",
			},
			// Add more fields here based on the API response and your needs
		},
	}
}
