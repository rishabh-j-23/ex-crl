package workflow

import (
	"fmt"
	"log"

	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a request from the workflow",
	Run: func(cmd *cobra.Command, args []string) {
		var wf models.Workflow
		err := utils.LoadJSONFile(utils.GetWorkflowFile(), &wf)
		if err != nil {
			log.Fatalf("Failed to load workflow config: %v", err)
		}

		if len(wf.Workflow) == 0 {
			fmt.Println("Workflow is empty.")
			return
		}

		// Build list of request names to pass to fzf
		var options []string
		for _, step := range wf.Workflow {
			options = append(options, step.RequestName)
		}

		// Show fzf to select items to remove (with --multi support)
		selected := utils.FzfFromList(options, true)
		if len(selected) == 0 {
			fmt.Println("No request selected for removal.")
			return
		}

		// Build a map for quick lookup
		selectedSet := make(map[string]bool)
		for _, sel := range selected {
			selectedSet[sel] = true
		}

		// Filter out the selected items
		var updated []models.WorkflowStep
		for _, step := range wf.Workflow {
			if !selectedSet[step.RequestName] {
				updated = append(updated, step)
			}
		}

		// Save updated workflow
		wf.Workflow = updated
		utils.SaveWorkflowConfig(wf)

		fmt.Println("Selected requests removed from workflow.")
	},
}

func init() {
	WorkflowCmd.AddCommand(removeCmd)
}
