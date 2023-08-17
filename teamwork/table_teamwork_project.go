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
			{
				Name:        "category_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the category.",
				Transform:   transform.FromField("Category.ID"),
			},
			{
				Name:        "category_color",
				Type:        proto.ColumnType_STRING,
				Description: "The color of the category.",
				Transform:   transform.FromField("Category.Color"),
			},
			{
				Name:        "category_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the category.",
				Transform:   transform.FromField("Category.Name"),
			},
			{
				Name:        "overview_start_page",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the Project's start page.",
				Transform:   transform.FromField("OverviewStartPage"),
			},
			{
				Name:        "integrations_xero_basecurrency",
				Type:        proto.ColumnType_STRING,
				Description: "The base currency used for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.Basecurrency"),
			},
			{
				Name:        "integrations_xero_countrycode",
				Type:        proto.ColumnType_STRING,
				Description: "The country code used for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.Countrycode"),
			},
			{
				Name:        "integrations_xero_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the Xero integration is enabled.",
				Transform:   transform.FromField("Integrations.Xero.Enabled"),
			},
			{
				Name:        "integrations_xero_connected",
				Type:        proto.ColumnType_STRING,
				Description: "The connected flag for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.Connected"),
			},
			{
				Name:        "integrations_xero_organisation",
				Type:        proto.ColumnType_STRING,
				Description: "The organisation used for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.organisation"),
			},
			{
				Name:        "integrations_sharepoint_account",
				Type:        proto.ColumnType_STRING,
				Description: "The account for Sharepoint integration.",
				Transform:   transform.FromField("Integrations.Sharepoint.Account"),
			},
			{
				Name:        "integrations_sharepoint_foldername",
				Type:        proto.ColumnType_STRING,
				Description: "The folder name for Sharepoint integration.",
				Transform:   transform.FromField("Integrations.Sharepoint.Foldername"),
			},
			{
				Name:        "integrations_sharepoint_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the Sharepoint integration is enabled.",
				Transform:   transform.FromField("Integrations.Sharepoint.Enabled"),
			},
			{
				Name:        "integrations_sharepoint_folder",
				Type:        proto.ColumnType_STRING,
				Description: "The folder for Sharepoint integration.",
				Transform:   transform.FromField("Integrations.Sharepoint.Folder"),
			},
			{
				Name:        "integrations_microsoftconnectors_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the Microsoft Connectors integration is enabled.",
				Transform:   transform.FromField("Integrations.MicrosoftConnectors.Enabled"),
			},
			{
				Name:        "integrations_onedrivebusiness_account",
				Type:        proto.ColumnType_STRING,
				Description: "The account for OneDrive Business integration.",
				Transform:   transform.FromField("Integrations.Onedrivebusiness.Account"),
			},
			{
				Name:        "integrations_onedrivebusiness_foldername",
				Type:        proto.ColumnType_STRING,
				Description: "The folder name for OneDrive Business integration.",
				Transform:   transform.FromField("Integrations.Onedrivebusiness.Foldername"),
			},
			{
				Name:        "integrations_onedrivebusiness_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the OneDrive Business integration is enabled.",
				Transform:   transform.FromField("Integrations.Onedrivebusiness.Enabled"),
			},
			{
				Name:        "integrations_onedrivebusiness_folder",
				Type:        proto.ColumnType_STRING,
				Description: "The folder  for OneDrive Business integration.",
				Transform:   transform.FromField("Integrations.Onedrivebusiness.Folder"),
			},
			{
				Name:        "defaults_privacy",
				Type:        proto.ColumnType_STRING,
				Description: "Privacy default information.",
				Transform:   transform.FromField("Defaults.Privacy"),
			},
			{
				Name:        "notify_everyone",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not to notify all project participants of changes.",
				Transform:   transform.FromField("Notifyeveryone"),
			},
			{
				Name:        "files_auto_new_version",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not file changes result in automatic new versions.",
				Transform:   transform.FromField("FilesAutoNewVersion"),
			},
			{
				Name:        "default_privacy",
				Type:        proto.ColumnType_STRING,
				Description: "Default privacy.",
				Transform:   transform.FromField("DefaultPrivacy"),
			},
			{
				Name:        "tasks_start_page",
				Type:        proto.ColumnType_STRING,
				Description: "URL of the tasks associated with this project.",
				Transform:   transform.FromField("TasksStartPage"),
			},
			{
				Name:        "starred",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the project is starred.",
				Transform:   transform.FromField("Starred"),
			},
			{
				Name:        "announcement_html",
				Type:        proto.ColumnType_STRING,
				Description: "HTML version of announcements associated with this project.",
				Transform:   transform.FromField("AnnouncementHTML"),
			},
			{
				Name:        "is_project_admin",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the calling user is an administrator of the project.",
				Transform:   transform.FromField("IsProjectAdmin"),
			},
			{
				Name:        "company_is_owner",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the calling user is an owner of the project's associated company.",
				Transform:   transform.FromField("Company.IsOwner"),
			},
			{
				Name:        "company_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Teamwork ID of the company associated with the project.",
				Transform:   transform.FromField("Company.ID"),
			},
			{
				Name:        "company_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the company associated with the project.",
				Transform:   transform.FromField("Company.Name"),
			},
			{
				Name:        "end_date",
				Type:        proto.ColumnType_STRING,
				Description: "The projected end date of the project.",
				Transform:   transform.FromField("EndDate"),
			},
			{
				Name:        "announcement",
				Type:        proto.ColumnType_STRING,
				Description: "Any announcements associated with the project.",
				Transform:   transform.FromField("Announcement"),
			},
			{
				Name:        "show_announcement",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not to show announcements associated with this project.",
				Transform:   transform.FromField("ShowAnnouncement"),
			},
			{
				Name:        "sub_status",
				Type:        proto.ColumnType_STRING,
				Description: "The project's sub-status.",
				Transform:   transform.FromField("SubStatus"),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: "Indicates whether or not to show announcements associated with this project.",
				Transform:   transform.FromField("Tags"),
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
