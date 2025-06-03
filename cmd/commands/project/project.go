package project

import (
	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/core"
	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

var editProjectFile bool

// projectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "perform project related operations",
	Long: `project is auto selected based on current dir. 
It is recommended to be in the root of project directory`,
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetString("set-env")
		assert.ErrIsNil(err, "Error parsing the val for --set-env")

		if flagVal != "" {
			core.SwitchEnv(flagVal)
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
