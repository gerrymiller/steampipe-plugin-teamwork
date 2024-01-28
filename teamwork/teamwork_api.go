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

/*
Sample usage:

	func main() {
		memberMap1 := map[string]interface{}{
			"Testing1": 1,
			"Testing2": "Cool!",
		}

		memberMap2 := map[string]interface{}{
			"Testing1": 2,
			"Testing2": "As beans!",
		}

		src1 := &Source{
			Name: "Sample 1",
			Members: []interface{}{
				10, 20, 30, memberMap1,
			},
		}

		src2 := &Source{
			Name: "Sample 2",
			Members: []interface{}{
				40, 50, 60, memberMap2,
			},
		}

		tgt := &Target{}

		if err := appendFirstArrayElement(src1, tgt); err != nil {
			panic(err)
		}

		if err := appendFirstArrayElement(src2, tgt); err != nil {
			panic(err)
		}

		fmt.Printf("Source: %+v\n", src1)
		fmt.Printf("Source: %+v\n", src2)
		fmt.Printf("Target: %+v\n", tgt)
	}

func appendFirstArrayElement(src, tgt interface{}) error {
	srcVal := reflect.ValueOf(src).Elem()
	tgtVal := reflect.ValueOf(tgt).Elem()

	var srcSlice reflect.Value
	var tgtSlice reflect.Value

	// Find the first slice field in the source
	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Field(i)
		if field.Kind() == reflect.Slice {
			srcSlice = field
			break
		}
	}

	// Find the first slice field in the target
	for i := 0; i < tgtVal.NumField(); i++ {
		field := tgtVal.Field(i)
		if field.Kind() == reflect.Slice {
			tgtSlice = field
			break
		}
	}

	if !srcSlice.IsValid() || !tgtSlice.IsValid() {
		return fmt.Errorf("could not find slice fields in both src and tgt")
	}

	// Append all elements from srcSlice to tgtSlice
	for i := 0; i < srcSlice.Len(); i++ {
		tgtSlice.Set(reflect.Append(tgtSlice, srcSlice.Index(i)))
	}

	return nil
}
*/

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

	/*
		if totalPages == -1 {
			totalPagesHeader := resp.Header.Get("x-pages")
			if totalPagesHeader != "" {
				totalPages, err = strconv.Atoi(resp.Header.Get("x-pages"))
				if err != nil {
					logger.Error(fmt.Sprintf("ListTeamworkItems(): Error %s", err.Error()))
					return nil, err
				}
			} else {
				totalPages = 1
			}
		}
	*/

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
