package workflow

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/rishabh-j-23/ex-crl/internal/models"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a request to the workflow",
	Run: func(cmd *cobra.Command, args []string) {
		// Load existing workflow config
		var wf models.Workflow
		err := utils.LoadJSONFile(utils.GetWorkflowFile(), &wf)
		if err != nil {
			slog.Error("Failed to load workflow config", "err", err)
			os.Exit(1)
		}

		// Select a request file using fzf
		requestDir := utils.GetRequestsDir()
		selected := utils.FzfSearch(requestDir)
		if selected == "" {
			fmt.Println("No request selected.")
			return
		}

		// Strip `.json` extension if present
		requestName := strings.TrimSuffix(filepath.Base(selected), filepath.Ext(selected))

		// Add the new workflow step
		step := models.WorkflowStep{
			RequestName: requestName,
			Exec:        true,
		}
		wf.Workflow = append(wf.Workflow, step)

		// Save updated workflow
		utils.SaveWorkflowConfig(wf)
		fmt.Printf("Added '%s' to workflow\n", requestName)
	},
}

func init() {
	WorkflowCmd.AddCommand(addCmd)
}
