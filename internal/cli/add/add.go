package add

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new request or query to your project.",
	Long: `Add a new HTTP request or query to your project.

Examples:
  ex-crl add request
  ex-crl add query
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
