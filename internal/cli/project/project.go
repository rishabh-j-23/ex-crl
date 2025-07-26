package project

import (
	"github.com/rishabh-j-23/ex-crl/internal/app"
	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

var editProjectFile bool

// projectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage project configuration and environments.",
	Long: `Manage project configuration and environments.

Examples:
  ex-crl project --edit         # Edit the project config file
  ex-crl project --set-env dev  # Switch to the 'dev' environment

Project is auto-selected based on the current directory. It is recommended to be in the root of your project directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetString("set-env")
		assert.ErrIsNil(err, "Error parsing the val for --set-env")

		if flagVal != "" {
			app.SwitchEnv(flagVal)
		}

		if editProjectFile {
			editor.LaunchEditor(utils.GetProjectConfig())
		}

	},
}

func init() {
	ProjectCmd.Flags().StringP("set-env", "", "", "set active env for the propject")
	ProjectCmd.Flags().BoolVarP(&editProjectFile, "edit", "e", false, "edit the project config file")
}
