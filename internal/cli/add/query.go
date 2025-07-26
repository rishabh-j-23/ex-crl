package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO: add support for graphql query
var queryCmd = &cobra.Command{
	Use:   "query [query-name] [endpoint]",
	Short: "Add a graphql query",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("query called")
	},
}

func init() {
	AddCmd.AddCommand(queryCmd)
}
