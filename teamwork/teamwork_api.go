package teamwork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-hclog"
)

func ListTeamworkItems(apiKey string, xurl string, response interface{},
	logger hclog.Logger) (interface{}, error) {

	logger.Trace(`Entering ListTeamworkItems()`)

	// Parse the URL so we can add query parameters
	u, err := url.Parse(xurl)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): req: %+v", req))
	logger.Trace(fmt.Sprintf("ListTeamworkItems(): url: %s", req.URL.String()))

	// Make sure we authorize with our API key, and that we accept JSON
	req.Header.Add(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":x")),
	)
	req.Header.Add("Accept", "application/json")

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): req: %+v", req.Header["Authorization"]))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()
	logger.Trace(fmt.Sprintf("ListTeamworkItems(): %s", resp.Status))
	logger.Trace(fmt.Sprintf("ListTeamworkItems(): Heeders: %+v", resp.Header))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
		return nil, err
	}

	// Execute pagination
	//resp.Header.Get()

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): body: %s", body))

	if err := json.Unmarshal(body, response); err != nil {
		logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
		return nil, err
	}

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): response: %+v", response))
	logger.Trace(`Exiting ListTeamworkItems()`)

	return response, nil
}
