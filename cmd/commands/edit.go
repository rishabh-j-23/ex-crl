package commands

import (
	"log"
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/editor"
	"github.com/rishabh-j-23/ex-crl/utils"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [name]",
	Short: "edit the json file for request and configs",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileToEdit := filepath.Join(utils.GetProjectDir(), utils.SelectFile(args, utils.GetProjectDir()))
		editor.LaunchEditor(fileToEdit)
	},
}

func init() {
	RootCmd.AddCommand(editCmd)
}
