package core

import (
	"fmt"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
)

func SwitchEnv(envName string) {
	var projectConfig models.ProjectConfig

	projectConfigFile := utils.GetProjectConfig()
	err := utils.LoadJSONFile(projectConfigFile, &projectConfig)
	assert.ErrIsNil(err, "Error loading the project config")

	if projectConfig.ActiveEnv.Name == envName {
		fmt.Printf("Switched to '%s'\n", envName)
		return
	}

	env := GetEnvByName(envName, projectConfig)
	assert.NotNil(env, fmt.Sprintf("'%s' does not exists", envName))

	projectConfig.ActiveEnv = *env

	utils.SaveProjectConfig(utils.GetProjectDir(), projectConfig)
	fmt.Printf("Switched to env '%s'\n", env.Name)
}

func GetEnvByName(envName string, projectConfig models.ProjectConfig) *models.Environment {
	for _, env := range projectConfig.ConfiguredEnv {
		if env.Name == envName {
			return &env
		}
	}
	return nil
}
