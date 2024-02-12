package teamwork

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTeamworkProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "teamwork_project",
		Description: "Projects from Teamwork.com",
		/*  // TODO - Implement GetConfig
		Get: &plugin.GetConfig{
			Hydrate: getTeamworkProject,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "id",
					Require:    plugin.Required,
					CacheMatch: "exact",
				},
			},
		},
		*/
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
				Transform:   transform.FromField("ID").NullIfZero(),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the project.",
				Transform:   transform.FromField("Name").NullIfZero(),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the project.",
				Transform:   transform.FromField("Description").NullIfZero(),
			},
			{
				Name:        "start_date",
				Type:        proto.ColumnType_STRING,
				Description: "The start date of the project.",
				Transform:   transform.FromField("StartDate").NullIfZero(),
			},
			{
				Name:        "last_changed_on",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date of the last change to the project.",
				Transform:   transform.FromField("LastChangedOn").NullIfZero(),
			},
			{
				Name:        "logo",
				Type:        proto.ColumnType_STRING,
				Description: "A URL to the project's logo.",
				Transform:   transform.FromField("Logo").NullIfZero(),
			},
			{
				Name:        "logo_from_company",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the logo is from a company.",
				Transform:   transform.FromField("LogoFromCompany").NullIfZero(),
			},
			{
				Name:        "created_on",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date the project was created.",
				Transform:   transform.FromField("CreatedOn").NullIfZero(),
			},
			{
				Name:        "privacy_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether this is a private project.",
				Transform:   transform.FromField("PrivacyEnabled").NullIfZero(),
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The current status of the project.",
				Transform:   transform.FromField("Status").NullIfZero(),
			},
			{
				Name:        "board_data",
				Type:        proto.ColumnType_JSON,
				Description: "Board data for the project.",
				Transform:   transform.FromField("BoardData").NullIfZero(),
			},
			{
				Name:        "reply_by_email_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the project supports replies via email.",
				Transform:   transform.FromField("ReplyByEmailEnabled").NullIfZero(),
			},
			{
				Name:        "harvest_timers_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether the project supports Harvest timers.",
				Transform:   transform.FromField("HarvestTimersEnabled").NullIfZero(),
			},
			{
				Name:        "category_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the category.",
				Transform:   transform.FromField("Category.ID").NullIfZero(),
			},
			{
				Name:        "category_color",
				Type:        proto.ColumnType_STRING,
				Description: "The color of the category.",
				Transform:   transform.FromField("Category.Color").NullIfZero(),
			},
			{
				Name:        "category_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the category.",
				Transform:   transform.FromField("Category.Name").NullIfZero(),
			},
			{
				Name:        "category_parent_id",
				Type:        proto.ColumnType_STRING,
				Description: "The parent ID of the category.",
				Transform:   transform.FromField("Category.ParentID").NullIfZero(),
			},
			{
				Name:        "overview_start_page",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the Project's start page.",
				Transform:   transform.FromField("OverviewStartPage").NullIfZero(),
			},
			{
				Name:        "integrations_xero_basecurrency",
				Type:        proto.ColumnType_STRING,
				Description: "The base currency used for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.Basecurrency").NullIfZero(),
			},
			{
				Name:        "integrations_xero_countrycode",
				Type:        proto.ColumnType_STRING,
				Description: "The country code used for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.Countrycode").NullIfZero(),
			},
			{
				Name:        "integrations_xero_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the Xero integration is enabled.",
				Transform:   transform.FromField("Integrations.Xero.Enabled").NullIfZero(),
			},
			{
				Name:        "integrations_xero_connected",
				Type:        proto.ColumnType_STRING,
				Description: "The connected flag for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.Connected").NullIfZero(),
			},
			{
				Name:        "integrations_xero_organisation",
				Type:        proto.ColumnType_STRING,
				Description: "The organisation used for Xero integration.",
				Transform:   transform.FromField("Integrations.Xero.organisation").NullIfZero(),
			},
			{
				Name:        "integrations_sharepoint_account",
				Type:        proto.ColumnType_STRING,
				Description: "The account for Sharepoint integration.",
				Transform:   transform.FromField("Integrations.Sharepoint.Account").NullIfZero(),
			},
			{
				Name:        "integrations_sharepoint_foldername",
				Type:        proto.ColumnType_STRING,
				Description: "The folder name for Sharepoint integration.",
				Transform:   transform.FromField("Integrations.Sharepoint.Foldername").NullIfZero(),
			},
			{
				Name:        "integrations_sharepoint_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the Sharepoint integration is enabled.",
				Transform:   transform.FromField("Integrations.Sharepoint.Enabled").NullIfZero(),
			},
			{
				Name:        "integrations_sharepoint_folder",
				Type:        proto.ColumnType_STRING,
				Description: "The folder for Sharepoint integration.",
				Transform:   transform.FromField("Integrations.Sharepoint.Folder").NullIfZero(),
			},
			{
				Name:        "integrations_microsoftconnectors_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the Microsoft Connectors integration is enabled.",
				Transform: transform.FromField("Integrations.MicrosoftConnectors.Enabled").
					NullIfZero(),
			},
			{
				Name:        "integrations_onedrivebusiness_account",
				Type:        proto.ColumnType_STRING,
				Description: "The account for OneDrive Business integration.",
				Transform: transform.FromField("Integrations.Onedrivebusiness.Account").
					NullIfZero(),
			},
			{
				Name:        "integrations_onedrivebusiness_foldername",
				Type:        proto.ColumnType_STRING,
				Description: "The folder name for OneDrive Business integration.",
				Transform: transform.FromField("Integrations.Onedrivebusiness.Foldername").
					NullIfZero(),
			},
			{
				Name:        "integrations_onedrivebusiness_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the OneDrive Business integration is enabled.",
				Transform: transform.FromField("Integrations.Onedrivebusiness.Enabled").
					NullIfZero(),
			},
			{
				Name:        "integrations_onedrivebusiness_folder",
				Type:        proto.ColumnType_STRING,
				Description: "The folder  for OneDrive Business integration.",
				Transform: transform.FromField("Integrations.Onedrivebusiness.Folder").
					NullIfZero(),
			},
			{
				Name:        "defaults_privacy",
				Type:        proto.ColumnType_STRING,
				Description: "Privacy default information.",
				Transform:   transform.FromField("Defaults.Privacy").NullIfZero(),
			},
			{
				Name:        "notify_everyone",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not to notify all project participants of changes.",
				Transform:   transform.FromField("Notifyeveryone").NullIfZero(),
			},
			{
				Name:        "files_auto_new_version",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not file changes result in automatic new versions.",
				Transform:   transform.FromField("FilesAutoNewVersion").NullIfZero(),
			},
			{
				Name:        "default_privacy",
				Type:        proto.ColumnType_STRING,
				Description: "Default privacy.",
				Transform:   transform.FromField("DefaultPrivacy").NullIfZero(),
			},
			{
				Name:        "start_page",
				Type:        proto.ColumnType_STRING,
				Description: "Which start page to use for this project.",
				Transform:   transform.FromField("StartPage").NullIfZero(),
			},
			{
				Name:        "tasks_start_page",
				Type:        proto.ColumnType_STRING,
				Description: "URL of the tasks associated with this project.",
				Transform:   transform.FromField("TasksStartPage").NullIfZero(),
			},
			{
				Name:        "starred",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the project is starred.",
				Transform:   transform.FromField("Starred").NullIfZero(),
			},
			{
				Name:        "announcement_html",
				Type:        proto.ColumnType_STRING,
				Description: "HTML version of announcements associated with this project.",
				Transform:   transform.FromField("AnnouncementHTML").NullIfZero(),
			},
			{
				Name:        "is_project_admin",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the calling user is an administrator of the project.",
				Transform:   transform.FromField("IsProjectAdmin").NullIfZero(),
			},
			{
				Name:        "is_billable",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the project is billable.",
				Transform:   transform.FromField("IsBillable").NullIfZero(),
			},
			{
				Name:        "is_onboarding_project",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the project represents onboarding.",
				Transform:   transform.FromField("IsOnboardingProject").NullIfZero(),
			},
			{
				Name:        "is_sample_project",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the project is a sample project.",
				Transform:   transform.FromField("IsSampleProject").NullIfZero(),
			},
			{
				Name:        "company_is_owner",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not the calling user is an owner of the project's associated company.",
				Transform:   transform.FromField("Company.IsOwner").NullIfZero(),
			},
			{
				Name:        "company_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Teamwork ID of the company associated with the project.",
				Transform:   transform.FromField("Company.ID").NullIfZero(),
			},
			{
				Name:        "company_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the company associated with the project.",
				Transform:   transform.FromField("Company.Name").NullIfZero(),
			},
			{
				Name:        "end_date",
				Type:        proto.ColumnType_STRING,
				Description: "The projected end date of the project.",
				Transform:   transform.FromField("EndDate").NullIfZero(),
			},
			{
				Name:        "announcement",
				Type:        proto.ColumnType_STRING,
				Description: "An announcement associated with the project.",
				Transform:   transform.FromField("Announcement").NullIfZero(),
			},
			{
				Name:        "show_announcement",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether or not to show announcements associated with this project.",
				Transform:   transform.FromField("ShowAnnouncement").NullIfZero(),
			},
			{
				Name:        "sub_status",
				Type:        proto.ColumnType_STRING,
				Description: "The project's sub-status.",
				Transform:   transform.FromField("SubStatus").NullIfZero(),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: "Tags associated with this project.",
				Transform:   transform.FromField("Tags").NullIfZero(),
			},
			{
				Name:        "portfolio_boards",
				Type:        proto.ColumnType_JSON,
				Description: "Portfolio boards associated with this project.",
				Transform:   transform.FromField("PortfolioBoards").NullIfZero(),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the project.",
				Transform:   transform.FromField("Type").NullIfZero(),
			},
			{
				Name:        "active_pages_billing",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Billing page is active.",
				Transform:   transform.FromField("ActivePages.Billing").NullIfZero(),
			},
			{
				Name:        "active_pages_comments",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Comments page is active.",
				Transform:   transform.FromField("ActivePages.Comments").NullIfZero(),
			},
			{
				Name:        "active_pages_files",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Files page is active.",
				Transform:   transform.FromField("ActivePages.Files").NullIfZero(),
			},
			{
				Name:        "active_pages_links",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Links page is active.",
				Transform:   transform.FromField("ActivePages.Links").NullIfZero(),
			},
			{
				Name:        "active_pages_notebooks",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Notebooks page is active.",
				Transform:   transform.FromField("ActivePages.Notebooks").NullIfZero(),
			},
			{
				Name:        "active_pages_tasks",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Tasks page is active.",
				Transform:   transform.FromField("ActivePages.Tasks").NullIfZero(),
			},
			{
				Name:        "active_pages_time",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Time page is active.",
				Transform:   transform.FromField("ActivePages.Time").NullIfZero(),
			},
			{
				Name:        "active_pages_risk_register",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Risk Register page is active.",
				Transform:   transform.FromField("ActivePages.RiskRegister").NullIfZero(),
			},
			{
				Name:        "active_pages_milestones",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Milestones page is active.",
				Transform:   transform.FromField("ActivePages.Milestones").NullIfZero(),
			},
			{
				Name:        "active_pages_messages",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Messages page is active.",
				Transform:   transform.FromField("ActivePages.Messages").NullIfZero(),
			},
			{
				Name:        "active_pages_board",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Board page is active.",
				Transform:   transform.FromField("ActivePages.Board").NullIfZero(),
			},
			{
				Name:        "active_pages_proofs",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Proofs page is active.",
				Transform:   transform.FromField("ActivePages.Proofs").NullIfZero(),
			},
			{
				Name:        "active_pages_table",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Table page is active.",
				Transform:   transform.FromField("ActivePages.Table").NullIfZero(),
			},
			{
				Name:        "active_pages_forms",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Forms page is active.",
				Transform:   transform.FromField("ActivePages.Forms").NullIfZero(),
			},
			{
				Name:        "active_pages_gantt",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Gantt page is active.",
				Transform:   transform.FromField("ActivePages.Gantt").NullIfZero(),
			},
			{
				Name:        "active_pages_finance",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the Finance page is active.",
				Transform:   transform.FromField("ActivePages.Finance").NullIfZero(),
			},
			{
				Name:        "active_pages_list",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the List page is active.",
				Transform:   transform.FromField("ActivePages.List").NullIfZero(),
			},
			{
				Name:        "direct_file_uploads_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether direct file uploads are allowed.",
				Transform:   transform.FromField("DirectFileUploadsEnabled").NullIfZero(),
			},
			{
				Name:        "skip_weekends",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether weekends are work days for this project.",
				Transform:   transform.FromField("SkipWeekends").NullIfZero(),
			},
		},
	}
}

func getTeamworkProject(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {
	// Logic to connect to Teamwork API and get a single project

	plugin.Logger(ctx).Trace("Entering getTeamworkProject()")

	config := GetConfig(d.Connection)

	var project ProjectResponse

	url := fmt.Sprintf("https://teamwork.%s.com", *config.Domain)
	url = fmt.Sprintf("%s/projects/%s.json", url, d.EqualsQualString("id"))

	plugin.Logger(ctx).Trace(fmt.Sprintf("getTeamworkProject(): url: %s", url))

	_, err := ListTeamworkItems(*config.APIKey, url, &project, plugin.Logger(ctx))
	if err != nil {
		plugin.Logger(ctx).Error(err.Error())
		return nil, err
	}

	plugin.Logger(ctx).Info(fmt.Sprintf("getTeamworkProject(): project %+v", project))

	plugin.Logger(ctx).Trace("Exiting getTeamworkProject()")
	return project.Project, nil
}

func listTeamworkProjects(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {
	// Logic to connect to Teamwork API and get a list of projects

	plugin.Logger(ctx).Trace("Entering listTeamworkProjects()")

	config := GetConfig(d.Connection)

	var projects ProjectsResponse

	url := fmt.Sprintf("https://teamwork.%s.com", *config.Domain)
	url = fmt.Sprintf("%s/projects.json", url)

	plugin.Logger(ctx).Trace(fmt.Sprintf("listTeamworkProjects(): url: %s", url))

	_, err := ListTeamworkItems(*config.APIKey, url, &projects, plugin.Logger(ctx))
	if err != nil {
		plugin.Logger(ctx).Error(err.Error())
		return nil, err
	}

	plugin.Logger(ctx).Info(fmt.Sprintf("listTeamworkProjects(): projects %+v", projects))

	for _, t := range projects.Projects {
		d.StreamListItem(ctx, t)
	}

	plugin.Logger(ctx).Trace("Exiting listTeamworkProjects()")
	return nil, nil
}

type Project struct {
	Announcement     string `json:"announcement"`
	AnnouncementHTML string `json:"announcementHTML"`
	BoardData        struct {
	} `json:"boardData"`
	CreatedOn                time.Time `json:"created-on"`
	DefaultPrivacy           string    `json:"defaultPrivacy"`
	Description              string    `json:"description"`
	DirectFileUploadsEnabled bool      `json:"directFileUploadsEnabled"`
	EndDate                  string    `json:"endDate"`
	FilesAutoNewVersion      bool      `json:"filesAutoNewVersion"`
	HarvestTimersEnabled     bool      `json:"harvest-timers-enabled"`
	ID                       string    `json:"id"`
	IsBillable               bool      `json:"isBillable"`
	IsOnBoardingProject      bool      `json:"isOnBoardingProject"`
	IsProjectAdmin           bool      `json:"isProjectAdmin"`
	IsSampleProject          bool      `json:"isSampleProject"`
	LastChangedOn            time.Time `json:"last-changed-on"`
	Logo                     string    `json:"logo"`
	LogoFromCompany          bool      `json:"logoFromCompany"`
	Name                     string    `json:"name"`
	Notifyeveryone           bool      `json:"notifyeveryone"`
	OverviewStartPage        string    `json:"overview-start-page"`
	PortfolioBoards          []any     `json:"portfolioBoards"`
	PrivacyEnabled           bool      `json:"privacyEnabled"`
	ReplyByEmailEnabled      bool      `json:"replyByEmailEnabled"`
	ShowAnnouncement         bool      `json:"show-announcement"`
	SkipWeekends             bool      `json:"skipWeekends"`
	Starred                  bool      `json:"starred"`
	StartPage                string    `json:"start-page"`
	StartDate                string    `json:"startDate"`
	Status                   string    `json:"status"`
	SubStatus                string    `json:"subStatus"`
	Tags                     []any     `json:"tags"`
	TasksStartPage           string    `json:"tasks-start-page"`
	Type                     string    `json:"type"`
	ActivePages              struct {
		Billing      string `json:"billing"`
		Board        string `json:"board"`
		Comments     string `json:"comments"`
		Files        string `json:"files"`
		Finance      string `json:"finance"`
		Forms        string `json:"forms"`
		Gantt        string `json:"gantt"`
		Links        string `json:"links"`
		List         string `json:"list"`
		Messages     string `json:"messages"`
		Milestones   string `json:"milestones"`
		Notebooks    string `json:"notebooks"`
		Proofs       string `json:"proofs"`
		RiskRegister string `json:"riskRegister"`
		Table        string `json:"table"`
		Tasks        string `json:"tasks"`
		Time         string `json:"time"`
	} `json:"active-pages"`
	Category struct {
		Color    string `json:"color"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		ParentID string `json:"parentId"`
	} `json:"category"`
	Company struct {
		ID      string `json:"id"`
		IsOwner string `json:"is-owner"`
		Name    string `json:"name"`
	} `json:"company"`
	Defaults struct {
		Privacy string `json:"privacy"`
	} `json:"defaults"`
	Integrations struct {
		Onedrivebusiness struct {
			Account    string `json:"account"`
			Enabled    bool   `json:"enabled"`
			Folder     string `json:"folder"`
			Foldername string `json:"foldername"`
		} `json:"onedrivebusiness"`
		Sharepoint struct {
			Account    string `json:"account"`
			Enabled    bool   `json:"enabled"`
			Folder     string `json:"folder"`
			Foldername string `json:"foldername"`
		} `json:"sharepoint"`
		Xero struct {
			Basecurrency string `json:"basecurrency"`
			Connected    string `json:"connected"`
			Countrycode  string `json:"countrycode"`
			Enabled      bool   `json:"enabled"`
			Organisation string `json:"organisation"`
		} `json:"xero"`
	} `json:"integrations"`
}

type ProjectsResponse struct {
	Status   string    `json:"STATUS"`
	Projects []Project `json:"projects"`
}

type ProjectResponse struct {
	Status  string  `json:"STATUS"`
	Project Project `json:"project"`
}
