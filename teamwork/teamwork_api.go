package teamwork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Config struct {
	APIKey string
	Domain string
}

type SDK struct {
	apiKey  string
	client  *http.Client
	baseURL string
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

func GetProjects(apiKey string, baseUrl string, logger hclog.Logger) (*ProjectsResponse, error) {
	logger.Trace(`Entering GetProjects()`)
	var projects ProjectsResponse

	ListTeamworkItems(apiKey, baseUrl+"/projects.json", &projects, logger)

	logger.Trace(fmt.Sprintf("%+v", projects.Projects))
	logger.Trace(`Exiting GetProjects()`)

	return &projects, nil
}

func ListTeamworkItems(apiKey string, url string, response interface{}, logger hclog.Logger) (interface{}, error) {
	logger.Trace(`Entering ListTeamworkItems()`)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Trace(fmt.Sprintf("req: %+v", req))
	logger.Trace(fmt.Sprintf("url: %s", req.URL.String()))

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":x")))
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(body, response); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Trace(fmt.Sprintf("response: %+v", response))
	logger.Trace(`Exiting ListTeamworkItems()`)

	return response, nil
}
