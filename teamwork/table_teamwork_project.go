package teamwork

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTeamworkProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "teamwork_project",
		Description: "Projects from Teamwork.com",
		List: &plugin.ListConfig{
			Hydrate: listTeamworkProjects,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the project.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the project.",
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the project.",
				Transform:   transform.FromField("Description"),
			},
			{
				Name:        "startDate",
				Type:        proto.ColumnType_STRING,
				Description: "The start date of the project",
				Transform:   transform.FromField("StartDate"),
			},
			{
				Name:        "lastChangedOn",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date of the last change to the project",
				Transform:   transform.FromField("LastChangedOn"),
			},
			{
				Name:        "logo",
				Type:        proto.ColumnType_STRING,
				Description: "A URL to the project's logo",
				Transform:   transform.FromField("Logo"),
			},
			{
				Name:        "createdOn",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date the project was created",
				Transform:   transform.FromField("CreatedOn"),
			},
			{
				Name:        "privacyEnabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether this is a private project",
				Transform:   transform.FromField("PrivacyEnabled"),
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The current status of the project",
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "boardData",
				Type:        proto.ColumnType_JSON,
				Description: "Board data for the project",
				Transform:   transform.FromField("BoardData"),
			},
			{
				Name:        "replyByEmailEnabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the project supports replies via email",
				Transform:   transform.FromField("ReplyByEmailEnabled"),
			},
			{
				Name:        "harvestTimersEnabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the project supports Harvest timers",
				Transform:   transform.FromField("HarvestTimersEnabled"),
			},
		},
	}
}

func listTeamworkProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Logic to connect to Teamwork API and get a list of projects

	config := GetConfig(d.Connection)

	projects, err := GetProjects(*config.APIKey, "https://teamwork."+*config.Domain+".com", plugin.Logger(ctx))
	if err != nil {
		return nil, err
	}

	for _, t := range (*projects).Projects {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}
