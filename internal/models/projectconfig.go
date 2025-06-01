package models

type ProjectConfig struct {
	Name          string            `json:"project"`
	ActiveEnv     Environment       `json:"active-env"`
	ConfiguredEnv []Environment     `json:"configured-env"`
	GlobalHeaders map[string]string `json:"global-headers"`
}
