package commands

import (
	"github.com/rishabh-j-23/ex-crl/internal/core"
	"os"

	"github.com/rishabh-j-23/ex-crl/cmd/commands/add"
	"github.com/rishabh-j-23/ex-crl/cmd/commands/project"
	"github.com/rishabh-j-23/ex-crl/cmd/commands/workflow"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ex-crl",
	Short: "A flexible and scriptable HTTP request CLI tool",
	Long: `ex-crl is a powerful CLI tool for managing and executing HTTP requests 
with support for workflows, environments, and reusable configurations.

Features include:
  - Defining requests using JSON files
  - Storing requests per project with environment-specific configs
  - Executing chained requests using workflows
  - Global and per-request headers

Use 'ex-crl project init' to nitialize a new project

Use "ex-crl [command] --help" for more information on a specific command.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		if err := core.InitLogger(debug); err != nil {
			// fallback to stderr if logger fails
			os.Stderr.WriteString("Failed to initialize logger: " + err.Error() + "\n")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// add a debuf flag
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
	RootCmd.AddCommand(add.AddCmd)
	RootCmd.AddCommand(project.ProjectCmd)
	RootCmd.AddCommand(workflow.WorkflowCmd)
	RootCmd.AddCommand(CompletionCmd)
}
