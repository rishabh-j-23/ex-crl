package add

import (
	"github.com/rishabh-j-23/ex-crl/internal/assert"
	"github.com/rishabh-j-23/ex-crl/internal/core"
	"github.com/spf13/cobra"
)

// requestCmd represents the request command
var requestCmd = &cobra.Command{
	Use:   "request [http-method] [request-name] [endpoint]",
	Short: "Add a new REST request definition.",
	Long: `Add a new REST request definition to your project.

Examples:
  ex-crl add request GET get-users /api/users
  ex-crl add request POST create-user /api/users`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		httpMethod := args[0]
		requestName := args[1]
		endpoint := args[2]
		// call your core logic here
		assert.EnsureNotEmpty(map[string]string{
			"httpMethod":  httpMethod,
			"requestName": requestName,
			"endpoint":    endpoint,
		})
		core.AddRequest(httpMethod, requestName, endpoint)
	},
}

func init() {
	AddCmd.AddCommand(requestCmd)
}
