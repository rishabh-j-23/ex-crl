package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/rishabh-j-23/ex-crl/utils/headers"
)

// InitProject initializes a new project with optional custom name, environment, and base URL.
func InitProject(projectName, envName, baseURL string) {
	if projectName == "" {
		projectName = utils.GetCurrentProjectName()
	}

	// Ensure main config directory exists
	utils.EnsureConfigDir(utils.ConfigDir)

	projectDir := utils.GetProjectDir()
	requestsDir := utils.GetRequestsDir()

	createDirs(projectDir, requestsDir)

	// Setup environment
	env := models.Environment{
		Name:    envName,
		BaseUrl: baseURL,
	}

	// Default global headers
	defaultHeaders := map[string]string{
		headers.ContentType: "application/json",
	}

	projectCfg := models.ProjectConfig{
		Name:          projectName,
		ActiveEnv:     env,
		ConfiguredEnv: []models.Environment{env},
		GlobalHeaders: defaultHeaders,
	}

	createJSONFileIfAbsent(projectDir, utils.ProjectConfigJson, projectCfg)
	utils.CreateWorflowFile()

	fmt.Printf("Project '%s' initialized at: %s\n", projectName, projectDir)
}

func createDirs(dirs ...string) {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			assert.ErrIsNil(err, fmt.Sprintf("❌ Failed to create directory: %s", dir))
		}
	}
}

// createJSONFileIfAbsent creates a file with marshalled JSON data if it does not already exist.
func createJSONFileIfAbsent(dir, filename string, data any) {
	filePath := filepath.Join(dir, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		utils.SaveDataToFile(filePath, data)
	}
}
