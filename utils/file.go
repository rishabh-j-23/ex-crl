package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/models"
)

func SaveDataToFile[T any](filePath string, data T) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	assert.ErrIsNil(err, "JSON marshaling failed for "+filePath)

	err = os.WriteFile(filePath, bytes, 0644)
	assert.ErrIsNil(err, "Failed to write file "+filePath)
}

// createJSONFileIfAbsent creates a file with marshalled JSON data if it does not already exist.
func CreateJSONFileIfAbsent(dir, filename string, data any) {
	filePath := filepath.Join(dir, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		SaveDataToFile(filePath, data)
	}
}

// Load the json file in the T
func LoadJSONFile[T any](path string, out *T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, out)
}

func GetFile(dir, requestName string) string {
	requestFile := filepath.Join(dir, requestName)

	_, err := os.Stat(requestFile)
	assert.ErrIsNil(err, fmt.Sprintf("%s request does not exists", requestName))

	return requestFile
}

func GetProjectConfig() string {
	if override := os.Getenv("EX_CRL_PROJECT_CONFIG"); override != "" {
		return override
	}
	projectsDir := GetProjectDir()
	return GetFile(projectsDir, ProjectConfigJson)
}

// Only returns file name with .json
func SelectFile(args []string, dir string) string {
	if len(args) == 0 {
		requestsDir := dir
		return FzfSearch(requestsDir)
	} else {
		assert.EnsureNotEmpty(map[string]string{
			"request-name": args[0],
		})
		return args[0] + ".json"
	}
}

func SaveProjectConfig(projectDir string, config models.ProjectConfig) {
	path := filepath.Join(projectDir, ProjectConfigJson)
	SaveDataToFile(path, config)
}

func GetWorkflowFile() string {
	return filepath.Join(GetProjectDir(), WorkflowConfigJson)
}

func SaveWorkflowConfig(config models.Workflow) {
	SaveDataToFile(GetWorkflowFile(), config)
}

func CreateWorflowFile() {
	workflowConfig := models.Workflow{
		Workflow: []models.WorkflowStep{
			{
				RequestName: "sample-request-name",
				Exec:        false,
			},
		},
	}
	fmt.Println("Creating workflow file")
	SaveWorkflowConfig(workflowConfig)
}
