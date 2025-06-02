package commands

import (
	"github.com/rishabh-j-23/ex-crl/internal/core"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [request-name]",
	Short: "Execute a stored HTTP request",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		requestName := utils.SelectFile(args, utils.GetRequestsDir())
		core.ExecRequest(requestName)
	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}
