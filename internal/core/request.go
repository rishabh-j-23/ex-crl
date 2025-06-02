package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
)

func AddRequest(httpMethod, requestName, endpoint string) {
	requestsDir := utils.GetRequestsDir()
	assert.EnsureDirExists(requestsDir)

	filePath := filepath.Join(requestsDir, requestName+".json")

	// Prevent overwrite if file exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("Request '%s' already exists\n", requestName)
		os.Exit(1)
	}

	request := models.Request{
		Name:       requestName,
		HttpMethod: strings.ToUpper(httpMethod),
		Headers:    map[string]string{},
		Body:       map[string]any{},
		Endpoint:   endpoint,
	}

	data, err := json.MarshalIndent(request, "", "  ")
	assert.ErrIsNil(err, "Failed to marshal request")

	err = os.WriteFile(filePath, data, 0644)
	assert.ErrIsNil(err, "Failed to write request file")

	editor.LaunchEditor(filePath)

	fmt.Printf("Request '%s' added at %s\n", requestName, filePath)
}
