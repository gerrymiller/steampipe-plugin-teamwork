package teamwork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

func fetchPage(apiKey, xurl string, page int) (*http.Response, error) {
	u, err := url.Parse(xurl)
	if err != nil {
		return nil, err
	}

	// Add 'page' query parameter
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":x")))
	req.Header.Add("Accept", "application/json")

	return http.DefaultClient.Do(req)
}

/*
Here's how I plan to implement this:
 1. Ensure that T is a struct
 2. Query T for a member called Status
 3. Set the value of Status to the value of the Status field in the response
 4. Query T for the next member after Status
 5. Determine if that member is a struct or a slice
 6. If a struct, set the value directly from the response
 7. If a slice, use pagination to append the values from the response calls to the slice
*/
func ListTeamworkItems[T any](
	apiKey, xurl string,
	response *T,
	logger hclog.Logger,
) (*T, error) {
	logger.Trace(`Entering ListTeamworkItems()`)

	responseValuePtr := reflect.ValueOf(response)
	if responseValuePtr.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("response must be a pointer to a struct")
	}

	responseStruct := responseValuePtr.Elem()
	if responseStruct.Kind() != reflect.Struct {
		return nil, fmt.Errorf("response must be a pointer to a struct")
	}

	statusField := responseStruct.FieldByName("Status")
	if !statusField.IsValid() {
		return nil, fmt.Errorf("response struct must have a field called Status")
	}
	if !statusField.CanSet() || statusField.Kind() != reflect.String {
		return nil, fmt.Errorf("status field must be an exported string field")
	}

	page := 1

	for {
		resp, err := fetchPage(apiKey, xurl, page)
		if err != nil {
			logger.Error(fmt.Sprintf("Error fetching page %d: %s", page, err))
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("Error reading response body: %s", err))
			return nil, err
		}

		var apiResponse T
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			logger.Error(fmt.Sprintf("Error unmarshalling response: %s", err))
			return nil, err
		}

		apiResponseValue := reflect.ValueOf(apiResponse)
		if apiResponseValue.Kind() != reflect.Struct {
			return nil, fmt.Errorf("API response must be a struct")
		}

		// For the first page, set the status. Assume status doesn't change across pages
		// TODO: maybe the status does change? Retry on error? Exponential backoff?
		if page == 1 {
			apiResponseStatusField := apiResponseValue.FieldByName("Status")
			if !apiResponseStatusField.IsValid() ||
				apiResponseStatusField.Kind() != reflect.String {
				return nil, fmt.Errorf("API response struct must have a field called Status")
			}

			statusField.SetString(apiResponseStatusField.String())
		}

		// Append items from the current page

		// Start by finding the 2nd field (assumes there are only two fields, Status and the data)
		dataField := responseStruct.Field(1)
		if dataField.Kind() == reflect.Slice && dataField.CanSet() {
			responseSlice := dataField.Interface()
			apiResponseSlice := apiResponseValue.Field(1).Interface()

			responseSliceValue := reflect.ValueOf(responseSlice)
			apiResponseSliceValue := reflect.ValueOf(apiResponseSlice)

			dataField.Set(reflect.AppendSlice(responseSliceValue, apiResponseSliceValue))
		}

		xPages := resp.Header.Get("x-pages")
		totalPages, _ := strconv.Atoi(xPages)

		if page >= totalPages {
			break
		}
		page++
	}

	logger.Trace(`Exiting ListTeamworkItems()`)
	return response, nil
}

/*
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
*/
