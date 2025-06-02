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
		}
		core.ExecRequest(requestName)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
