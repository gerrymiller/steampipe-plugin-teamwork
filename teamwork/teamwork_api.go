package teamwork

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

// fetchPage is a helper function to fetch data from the API.
func fetchPage(apiKey, url string, page int) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", url, page), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	return client.Do(req)
}

// unmarshalResponse unmarshals the http response body into the given struct pointer.
func unmarshalResponse(resp *http.Response, target interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

// setPageData handles setting data from API response to response struct.
func setPageData(apiResponse, responseData reflect.Value, page int) error {
	if page == 1 {
		status := apiResponse.FieldByName("Status")
		if !status.IsValid() || status.Kind() != reflect.String {
			return fmt.Errorf("API response struct must have a field called Status")
		}
		responseData.FieldByName("Status").SetString(status.String())
	}

	dataField := responseData.Field(1)
	apiDataField := apiResponse.Field(1)

	if dataField.Kind() == reflect.Slice && dataField.CanSet() {
		dataField.Set(reflect.AppendSlice(dataField, apiDataField))
	} else if dataField.Kind() == reflect.Struct && dataField.CanSet() {
		dataField.Set(apiDataField)
	} else {
		return fmt.Errorf("unsupported data field type or field not settable")
	}
	return nil
}

// ListTeamworkItems fetches items from teamwork API and populates them into the given response struct.
func ListTeamworkItems[T any](apiKey, url string, response *T, logger hclog.Logger) (*T, error) {
	logger.Trace("Entering ListTeamworkItems()")
	defer logger.Trace("Exiting ListTeamworkItems()")

	if reflect.TypeOf(response).Kind() != reflect.Ptr ||
		reflect.ValueOf(response).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("response must be a pointer to a struct")
	}

	responseData := reflect.ValueOf(response).Elem()
	page, totalPages := 1, 1

	for page <= totalPages {
		resp, err := fetchPage(apiKey, url, page)
		if err != nil {
			logger.Error(fmt.Sprintf("Error fetching page %d: %s", page, err))
			return nil, err
		}
		defer resp.Body.Close()

		var apiResponse T
		if err := unmarshalResponse(resp, &apiResponse); err != nil {
			logger.Error(fmt.Sprintf("Error unmarshalling response: %s", err))
			return nil, err
		}

		if err := setPageData(reflect.ValueOf(apiResponse), responseData, page); err != nil {
			return nil, err
		}

		if page == 1 { // Only read total pages once
			if xPages := resp.Header.Get("x-pages"); xPages != "" {
				totalPages, _ = strconv.Atoi(xPages)
			}
		}
		page++
	}
	return response, nil
}
