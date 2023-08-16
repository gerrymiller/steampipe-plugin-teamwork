package teamwork

import (
	"context"
	"time"

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
				Name:        "start_date",
				Type:        proto.ColumnType_STRING,
				Description: "The start date of the project.",
				Transform:   transform.FromField("StartDate"),
			},
			{
				Name:        "last_changed_on",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date of the last change to the project.",
				Transform:   transform.FromField("LastChangedOn"),
			},
			{
				Name:        "logo",
				Type:        proto.ColumnType_STRING,
				Description: "A URL to the project's logo.",
				Transform:   transform.FromField("Logo"),
			},
			{
				Name:        "created_on",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date the project was created.",
				Transform:   transform.FromField("CreatedOn"),
			},
			{
				Name:        "privacy_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether this is a private project.",
				Transform:   transform.FromField("PrivacyEnabled"),
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The current status of the project.",
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "board_data",
				Type:        proto.ColumnType_JSON,
				Description: "Board data for the project.",
				Transform:   transform.FromField("BoardData"),
			},
			{
				Name:        "reply_by_email_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the project supports replies via email.",
				Transform:   transform.FromField("ReplyByEmailEnabled"),
			},
			{
				Name:        "harvest_timers_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the project supports Harvest timers.",
				Transform:   transform.FromField("HarvestTimersEnabled"),
			},
		},
	}
}

func listTeamworkProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Logic to connect to Teamwork API and get a list of projects

	plugin.Logger(ctx).Trace("Entering listTeamworkProjects()")

	config := GetConfig(d.Connection)

	var projects ProjectsResponse
	_, err := ListTeamworkItems(*config.APIKey, "https://teamwork."+*config.Domain+".com/projects.json", &projects, plugin.Logger(ctx))
	if err != nil {
		plugin.Logger(ctx).Error(err.Error())
		return nil, err
	}

	for _, t := range projects.Projects {
		d.StreamListItem(ctx, t)
	}

	plugin.Logger(ctx).Trace("Exiting listTeamworkProjects()")
	return nil, nil
}

type ProjectsResponse struct {
	Status   string `json:"STATUS"`
	Projects []struct {
		StartDate      string    `json:"startDate"`
		LastChangedOn  time.Time `json:"last-changed-on"`
		Logo           string    `json:"logo"`
		CreatedOn      time.Time `json:"created-on"`
		PrivacyEnabled bool      `json:"privacyEnabled"`
		Status         string    `json:"status"`
		BoardData      struct {
		} `json:"boardData"`
		ReplyByEmailEnabled  bool   `json:"replyByEmailEnabled"`
		HarvestTimersEnabled bool   `json:"harvest-timers-enabled"`
		Description          string `json:"description"`
		Category             struct {
			Color string `json:"color"`
			ID    string `json:"id"`
			Name  string `json:"name"`
		} `json:"category"`
		ID                string `json:"id"`
		OverviewStartPage string `json:"overview-start-page"`
		StartPage         string `json:"start-page"`
		Integrations      struct {
			Xero struct {
				Basecurrency string `json:"basecurrency"`
				Countrycode  string `json:"countrycode"`
				Enabled      bool   `json:"enabled"`
				Connected    string `json:"connected"`
				Organisation string `json:"organisation"`
			} `json:"xero"`
			Sharepoint struct {
				Account    string `json:"account"`
				Foldername string `json:"foldername"`
				Enabled    bool   `json:"enabled"`
				Folder     string `json:"folder"`
			} `json:"sharepoint"`
			MicrosoftConnectors struct {
				Enabled bool `json:"enabled"`
			} `json:"microsoftConnectors"`
			Onedrivebusiness struct {
				Account    string `json:"account"`
				Foldername string `json:"foldername"`
				Enabled    bool   `json:"enabled"`
				Folder     string `json:"folder"`
			} `json:"onedrivebusiness"`
		} `json:"integrations"`
		Defaults struct {
			Privacy string `json:"privacy"`
		} `json:"defaults"`
		Notifyeveryone      bool   `json:"notifyeveryone"`
		FilesAutoNewVersion bool   `json:"filesAutoNewVersion"`
		DefaultPrivacy      string `json:"defaultPrivacy"`
		TasksStartPage      string `json:"tasks-start-page"`
		Starred             bool   `json:"starred"`
		AnnouncementHTML    string `json:"announcementHTML"`
		IsProjectAdmin      bool   `json:"isProjectAdmin"`
		Name                string `json:"name"`
		Company             struct {
			IsOwner string `json:"is-owner"`
			ID      string `json:"id"`
			Name    string `json:"name"`
		} `json:"company"`
		EndDate          string `json:"endDate"`
		Announcement     string `json:"announcement"`
		ShowAnnouncement bool   `json:"show-announcement"`
		SubStatus        string `json:"subStatus"`
		Tags             []any  `json:"tags"`
	} `json:"projects"`
}
