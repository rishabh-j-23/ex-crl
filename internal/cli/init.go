package cli

import (
	"github.com/rishabh-j-23/ex-crl/internal/app"
	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/spf13/cobra"
)

var projectName string
var envName string
var baseUrl string

var initCmd = &cobra.Command{
	Use:   "init [project-name] [environment] [base-api-url]",
	Short: "Initialize a new ex-crl project",
	Long:  "Initialize a new ex-crl project with optional project name, environment, and API base URL.",
	Args:  cobra.MaximumNArgs(3), // allow up to 3 args (all optional)
	Run: func(cmd *cobra.Command, args []string) {
		var projectName, envName, baseUrl string

		if len(args) > 0 {
			projectName = args[0]
		}
		if len(args) > 1 {
			envName = args[1]
		}
		if len(args) > 2 {
			baseUrl = args[2]
		}

		assert.EnsureNotEmpty(map[string]string{
			"envName": envName,
			"baseUrl": baseUrl,
		})

		app.InitProject(projectName, envName, baseUrl)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
