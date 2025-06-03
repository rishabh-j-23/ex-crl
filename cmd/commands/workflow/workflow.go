package workflow

import (
	"fmt"

	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

var editWorkflow bool

var WorkflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Manage workflows",
	Long:  `This command is used to view or edit the workflow file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if editWorkflow {
			editor.LaunchEditor(utils.GetWorkflowFile())
			return
		}
		fmt.Println("workflow called")
	},
}

func init() {
	WorkflowCmd.Flags().BoolVarP(&editWorkflow, "edit", "e", false, "Edit the workflow file")
}
