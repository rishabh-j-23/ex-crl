package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
)

const AppName = "ex-crl"
const ProjectConfigJson = "projectconfig.json"
const WorkflowConfigJson = "workflow.json"
const HeadersJson = "headers.json"

var ConfigDir string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home directory: %v", err)
	}
	ConfigDir = filepath.Join(home, AppName)

	EnsureConfigDir(ConfigDir)
}

func EnsureConfigDir(dir string) {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("failed to create directory %s: %v", dir, err)
		}
	} else if err != nil {
		log.Fatalf("error checking directory %s: %v", dir, err)
	} else if !info.IsDir() {
		log.Fatalf("%s exists but is not a directory", dir)
	}
}

func GetCurrentProjectName() string {
	currentPath, err := os.Getwd()
	assert.ErrIsNil(err, "Error getting the pwd")
	projectName := filepath.Base(currentPath)
	return projectName
}
