package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	httpReq, err := models.CreateRequest(request, projectConfig)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println(request.Name, "not executed succesfully")
		return
	}
	log.Println(request.Name, "executed succesfully")
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body) // ignore response
}
