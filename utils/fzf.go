package utils

import (
	"os"
	"os/exec"
	"strings"
)

// Only returns file name with .json
func FzfSearch(dir string) string {
	cmd := exec.Command(
		"fzf",
		"--height", "20", // Only take up 40% of the terminal height
		"--layout", "reverse", // Show results at the bottom
		"--preview", "bat --style=numbers --color=always --paging=never "+dir+"/{}",
		"--preview-window", "right:60%",
		"--style", "minimal",
	)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		println("Failed to select request:", err.Error())
		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}
