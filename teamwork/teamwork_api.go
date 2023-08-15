package teamwork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Add more fields as per Teamwork's API
}

func New(config Config) *SDK {
	return &SDK{
		apiKey:  base64.StdEncoding.EncodeToString([]byte(config.APIKey)),
		client:  &http.Client{},
		baseURL: fmt.Sprintf("https://teamwork.%s.com", config.Domain),
	}
}

func (sdk *SDK) GetProjects(logger hclog.Logger) ([]Project, error) {

	logger.Debug(`Entering GetProjects()`)
	logger.Debug(sdk.apiKey)

	req, err := http.NewRequest("GET", sdk.baseURL+"/projects.json", nil)
	logger.Debug(fmt.Sprintf("%+v", req))
	logger.Debug(req.URL.String())

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug(sdk.apiKey)

	req.Header.Add("Authorization", "Basic "+sdk.apiKey)
	req.Header.Add("Accept", "application/json")

	resp, err := sdk.client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var projects struct {
		Projects []Project `json:"projects"`
	}

	if err := json.Unmarshal(body, &projects); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug(fmt.Sprintf("%+v", projects.Projects))

	logger.Debug(`Exiting GetProjects()`)
	return projects.Projects, nil
}
