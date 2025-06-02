package commands

import (
	"github.com/rishabh-j-23/ex-crl/internal/assert"
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
		var requestName string
		if len(args) == 0 {
			requestsDir := utils.GetRequestsDir()
			requestName = utils.FzfSearch(requestsDir)
		} else {
			requestName = args[0]
			assert.EnsureNotEmpty(map[string]string{
				"request-name": requestName,
			})
			requestName = requestName + ".json"
		}
		core.ExecRequest(requestName)
	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}
