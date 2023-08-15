package teamwork

import (
	"context"
	"log"

	//	"fmt"
	"net/http"
	//	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const baseURL = "https://{{.Domain}}.teamwork.com"

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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Add more fields as per Teamwork's API
}

func GetTeamworkProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Call the listTeamworkProjects function to get a list of projects
	return listTeamworkProjects(ctx, d, nil)
}

func listTeamworkProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Logic to connect to Teamwork API and get a list of projects

	// For now, we'll return dummy data:
	dummyProject := map[string]interface{}{
		"name":        "Sample Project",
		"description": "A sample project description.",
	}
	log.Print(dummyProject)
	d.StreamListItem(ctx, dummyProject)

	return nil, nil
}
