package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
)

func InitProject(projectName string, envName string, baseUrl string) {
	// set project to basename of current dir
	if projectName == "" {
		projectName = utils.GetCurrentProjectName()
	}

	// create the dir to store all projects and configs and all
	utils.EnsureConfigDir(utils.ConfigDir)

	projectDir := filepath.Join(utils.ConfigDir, "projects", projectName)
	requestsDir := filepath.Join(projectDir, "requests")

	dirs := []string{
		projectDir,
		requestsDir,
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		assert.ErrIsNil(err, fmt.Sprintf("Failed to create directory %s", dir))
	}

	currentEnv := models.Environment{
		Name:    envName,
		BaseUrl: baseUrl,
	}

	// Add current environment to configured environments slice
	configuredEnvs := []models.Environment{currentEnv}

	// Full project config
	projectCfg := models.ProjectConfig{
		Name:          projectName,
		ActiveEnv:     currentEnv,
		ConfiguredEnv: configuredEnvs,
	}

	createFileWithStruct(projectDir, utils.ProjectConfigJson, projectCfg)

	fmt.Printf("Project '%s' initialized at %s\n", projectName, projectDir)
}

func createFileWithStruct(dir, filename string, data any) {
	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		bytes, err := json.MarshalIndent(data, "", "  ")
		assert.ErrIsNil(err, "Failed to marshal JSON for "+filename)
		err = os.WriteFile(path, bytes, 0644)
		assert.ErrIsNil(err, "Failed to write "+filename)
	}
}
