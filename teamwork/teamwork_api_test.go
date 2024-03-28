package teamwork

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
)

func setupSuite(tb testing.TB) (func(tb testing.TB), string) {
	tb.Log("setupSuite")

	// Create a test server that always returns the same response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var file string
		projectPattern := regexp.MustCompile(`^/project/[0-9]+\.json$`)
		projectsPattern := regexp.MustCompile(`^/projects(_paginated)?\.json$`)

		switch cmd := r.URL.Path; {
		// Return an unpaginated list of projects
		case projectsPattern.MatchString(cmd):
			file = `test_data/projects.json`
		// Return a single project
		case projectPattern.MatchString(cmd):
			file = `test_data/project.json`
		default:
			tb.Errorf("unexpected path: %v", r.URL.Path)
		}
		contents, err := os.ReadFile(file)
		if err != nil {
			tb.Errorf("unexpected error: %v", err)
		}
		w.Header().Set("Server", "nginx")
		w.Header().Set("Date", time.Now().In(time.FixedZone("GMT", 0)).Format(http.TimeFormat))
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Connection", "keep-alive")
		w.Header().
			Set("access-control-allow-headers", "Authorization,Content-Type,X-Set-WWW-Authenticate,X-Requested-With")
		w.Header().Set("access-control-allow-methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-expose-headers", "id,x-page,x-pages,x-records")
		w.Header().Set("access-control-max-age", "1000")
		w.Header().
			Set("cache-control", "private,must-revalidate,max-stale=0,max-age=0,post-check=0,pre-check=0")

		// Generate a random ETag value
		randomBytes := make([]byte, 16)
		_, err = rand.Read(randomBytes)
		if err != nil {
			tb.Errorf("Error generating random ETag: %v", err)
		}
		etag := base64.URLEncoding.EncodeToString(randomBytes)
		w.Header().Set("etag", etag)

		w.Header().Set("vary", "Origin")
		w.Header().Set("x-api-version", "region: 'us-east-1' env: 'prod' commit: '2185e2f'")
		w.Header().Set("x-content-type-options", "nosniff")
		w.Header().Set("x-from-cache", "true")
		w.Header().Set("x-isfiltered:", "false")
		w.Header().Set("x-lastupdated", "2024-01-28T14:43:33Z")

		page := r.URL.Query().Get("page")
		if page == "" {
			w.Header().Set("x-page", "1")
		} else {
			w.Header().Set("x-page", page)
		}

		if r.URL.Path == "/projects_paginated.json" {
			w.Header().Set("x-pages", "2")
			w.Header().Set("x-records", "142")
		} else {
			w.Header().Set("x-pages", "1")
			w.Header().Set("x-records", "71")
		}

		w.Header().Set("x-ratelimit-limit", "150")
		w.Header().Set("x-ratelimit-remaining", "149")
		w.Header().Set("x-ratelimit-reset", "60")
		w.Header().Set("x-xss-protection", "1; mode=block")
		w.Header().Set("access-control-allow-credentials", "true")
		w.Header().
			Set("content-security-policy", "frame-ancestors 'self' localhost *.teamwork.com *.teamworkpm.net teams.microsoft.com *.teams.microsoft.com *.skype.com teamworkintegrations.ngrok.io *.us.teamworkops.com;")

		w.Write(contents)
	}))

	return func(tb testing.TB) {
		tb.Log("teardownSuite")
		ts.Close()
	}, ts.URL
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	tb.Log("setupTest")

	return func(tb testing.TB) {
		tb.Log("teardownTest")
	}
}

func TestListTeamworkItemsProjects(t *testing.T) {

	teardownSuite, url := setupSuite(t)
	defer teardownSuite(t)

	for _, test := range []struct {
		Name string
		fn   func(*testing.T, string)
	}{
		//{"testListTeamworkItemsProject", testListTeamworkItemsProject},
		{"testListTeamworkItemsProjectsUnpaginated", testListTeamworkItemsProjectsUnpaginated},
		{"testListTeamworkItemsProjectsPaginated", testListTeamworkItemsProjectsPaginated},
	} {
		teardownTest := setupTest(t)
		defer teardownTest(t)
		t.Run(test.Name, func(t *testing.T) {
			test.fn(t, url)
		})
	}
}

/*
func testListTeamworkItemsProject(t *testing.T, url string) {
	// Call the API
	var response ProjectResponse
	_, err := ListTeamworkItems("apiKey", url+"/project/483331.json", &response, hclog.Default())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if response.Status != "OK" {
		t.Errorf("unexpected status: got %v, want %v", response.Status, "OK")
	}
	if response.Project.ID != "483331" {
		t.Errorf("unexpected project ID: got %v, want %v", response.Project.ID, "483331")
	}
}
*/

func testListTeamworkItemsProjectsUnpaginated(t *testing.T, url string) {
	// Call the API
	var response ProjectsResponse
	_, err := ListTeamworkItems("apiKey", url+"/projects.json", &response, hclog.Default())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if response.Status != "OK" {
		t.Errorf("unexpected status: got %v, want %v", response.Status, "OK")
	}
	if len(response.Projects) != 71 {
		t.Errorf("unexpected number of projects: got %v, want %v", len(response.Projects), 71)
	}
}

func testListTeamworkItemsProjectsPaginated(t *testing.T, url string) {
	// Call the API
	var response ProjectsResponse
	_, err := ListTeamworkItems(
		"apiKey",
		url+"/projects_paginated.json",
		&response,
		hclog.Default(),
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if response.Status != "OK" {
		t.Errorf("unexpected status: got %v, want %v", response.Status, "OK")
	}
	if len(response.Projects) != 142 {
		t.Errorf("unexpected number of projects: got %v, want %v", len(response.Projects), 142)
	}
}
