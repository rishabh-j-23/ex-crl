package commands

import (
	"log"
	"os"
	"path/filepath"
	"time"

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
		if cmd.Flag("debug").Changed {
			log.SetOutput(os.Stdout)
			log.SetFlags(log.LstdFlags | log.Lshortfile)
		} else {
			log.SetFlags(log.LstdFlags)

			logDir := filepath.Join(os.Getenv("HOME"), "ex-crl", "logs")
			logFileName := "ex-crl_" + time.Now().Format("2006-01-02") + ".log"
			logPath := filepath.Join(logDir, logFileName)

			// check if log directory exists, create it if not
			if _, err := os.Stat(logDir); os.IsNotExist(err) {
				err := os.MkdirAll(logDir, 0755)
				if err != nil {
					log.Printf("Failed to create log directory (%s): %v", logDir, err)
				}
			}

			logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				// fallback to stdout if file can't be opened
				log.SetOutput(os.Stdout)
				log.Printf("Failed to open log file (%s), logging to stdout: %v", logPath, err)
			} else {
				log.SetOutput(logFile)
			}
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
}
