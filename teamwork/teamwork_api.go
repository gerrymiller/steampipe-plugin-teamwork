package teamwork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

func ListTeamworkItems(apiKey string, url string, response interface{}, logger hclog.Logger) (interface{}, error) {
	logger.Trace(`Entering ListTeamworkItems()`)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): req: %+v", req))
	logger.Trace(fmt.Sprintf("ListTeamworkItems(): url: %s", req.URL.String()))

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":x")))
	req.Header.Add("Accept", "application/json")

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): req: %+v", req.Header["Authorization"]))

	resp, err := http.DefaultClient.Do(req)

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): %s", resp.Status))

	if err != nil {
		logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
		return nil, err
	}

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): body: %s", body))

	if err := json.Unmarshal(body, response); err != nil {
		logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
		return nil, err
	}

	logger.Trace(fmt.Sprintf("ListTeamworkItems(): response: %+v", response))
	logger.Trace(`Exiting ListTeamworkItems()`)

	return response, nil
}
