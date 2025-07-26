package cli

import (
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/app"
	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

var editBeforeExec bool

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [request-name]",
	Short: "Execute a stored HTTP request",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		requestName := utils.SelectFile(args, utils.GetRequestsDir())

		if editBeforeExec {
			editor.LaunchEditor(filepath.Join(utils.GetRequestsDir(), requestName))
		}

		app.ExecRequest(requestName)
	},
}

func init() {
	execCmd.Flags().BoolVarP(&editBeforeExec, "edit", "e", false, "Edit the request before executing it")
	RootCmd.AddCommand(execCmd)
}
