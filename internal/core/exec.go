package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
)

func ExecRequest(requestName string) {
	var request models.Request

	path := utils.GetFile(utils.GetRequestsDir(), requestName)
	err := utils.LoadJSONFile(path, &request)

	assert.ErrIsNil(err, fmt.Sprintf("Error unmarshalling json file %s", requestName))

	PerformRequest(request)
}

func PerformRequest(request models.Request) {
	var projectConfig models.ProjectConfig
	utils.LoadJSONFile(utils.GetProjectConfig(), &projectConfig)

	start := time.Now()
	httpReq, err := models.CreateRequest(request, projectConfig)
	assert.ErrIsNil(err, "Error creating the http request")

	// Example: Use it with http.Client
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	assert.ErrIsNil(err, "Error while performing the request")

	bodyBytes, err := io.ReadAll(resp.Body)
	assert.ErrIsNil(err, "Failed to read response body")

	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err != nil {
		// If it's not JSON, just print raw
		fmt.Println(string(bodyBytes))
	} else {
		fmt.Println(prettyJSON.String())
	}
	end := time.Now()
	fmt.Println("total time:", end.Sub(start))
}
