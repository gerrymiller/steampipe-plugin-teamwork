package teamwork

import (
	"context"
	"os"

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
				Name:        "ID",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the project.",
			},
			{
				Name:        "Name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the project.",
			},
			{
				Name:        "Description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the project.",
			},
			// Add more fields here based on the API response and your needs
		},
	}
}

func listTeamworkProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Logic to connect to Teamwork API and get a list of projects

	/*
		// For now, we'll return dummy data:
		dummyProject := map[string]interface{}{
			"name":        "Sample Project",
			"description": "A sample project description.",
		}
		d.StreamListItem(ctx, dummyProject)
	*/

	config := Config{
		APIKey: os.Getenv("TEAMWORK_API_KEY"),
		Domain: os.Getenv("TEAMWORK_DOMAIN"),
	}

	sdk := New(config)

	projects, err := sdk.GetProjects(plugin.Logger(ctx))
	if err != nil {
		return nil, nil
	}

	for _, t := range projects {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}
