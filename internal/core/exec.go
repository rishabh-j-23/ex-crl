package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
)

func ExecRequest(requestName string) {
	// First run the workflow steps if any
	runWorkflowSteps()

	// Then run the main request
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
	httpReq, err := CreateRequest(request, projectConfig)
	assert.ErrIsNil(err, "Error creating the HTTP request")

	client := &http.Client{
		Jar: utils.LoadCookiesFromDisk(), // Load persistent cookies
	}
	resp, err := client.Do(httpReq)
	assert.ErrIsNil(err, "Error while performing the request")

	utils.TrackDomain(httpReq.URL)
	utils.SaveCookiesToDisk()

	defer resp.Body.Close()

	end := time.Now()

	printRequestDetails(httpReq, resp, end.Sub(start))
}

func printRequestDetails(httpReq *http.Request, httpResp *http.Response, timeTakenByRequest time.Duration) {
	bodyBytes, err := io.ReadAll(httpResp.Body)
	assert.ErrIsNil(err, "Failed to read response body")

	fmt.Printf("%s %s %s %s\n", httpReq.Proto, httpReq.Method, httpReq.URL.RequestURI(), httpResp.Status)
	fmt.Printf("Host: %s\n\n", httpReq.URL.Hostname())

	// Pretty print JSON if possible
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err != nil {
		// If it's not JSON, just print raw
		fmt.Println(string(bodyBytes))
	} else {
		fmt.Println(prettyJSON.String())
	}

	fmt.Printf("\n[time taken: %v]\n", timeTakenByRequest)
}

func runWorkflowSteps() {
	var workflow models.Workflow
	err := utils.LoadJSONFile(utils.GetWorkflowFile(), &workflow)
	if err != nil {
		// No workflow file or load failed, skip silently
		return
	}

	for _, step := range workflow.Workflow {
		if step.Exec {
			var req models.Request
			path := utils.GetFile(utils.GetRequestsDir(), step.RequestName+".json")
			err := utils.LoadJSONFile(path, &req)
			if err == nil {
				executeSilently(req)
			}
		}
	}
}

func executeSilently(request models.Request) {
	var projectConfig models.ProjectConfig
	utils.LoadJSONFile(utils.GetProjectConfig(), &projectConfig)

	httpReq, err := CreateRequest(request, projectConfig)
	if err != nil {
		return
	}

	client := &http.Client{
		Jar: utils.LoadCookiesFromDisk(), // Load persistent cookies
	}
	resp, err := client.Do(httpReq)
	assert.ErrIsNil(err, "Error while performing the request")

	utils.TrackDomain(httpReq.URL)
	utils.SaveCookiesToDisk()

	slog.Info("Request executed successfully", "request", request.Name)
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body) // ignore response
}
