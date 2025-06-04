package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// FzfCmd returns the configured *exec.Cmd to run fzf with preview for files in dir
func FzfCmd(dir string) *exec.Cmd {
	return exec.Command(
		"fzf",
		"--height", "20",
		"--layout", "reverse",
		"--preview", "bat --style=numbers --color=always --paging=never "+dir+"/{}",
		"--preview-window", "right:60%",
		"--style", "minimal",
	)
}

// FzfSearch shows a file list in dir and returns selected file name (with .json suffix)
func FzfSearch(dir string) string {
	cmd := FzfCmd(dir)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = nil
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Cancelled or error in fzf:", err)
		os.Exit(0)
	}

	return strings.TrimSpace(string(out))
}

// FzfFromList shows a list of strings using fzf and returns selected items
func FzfFromList(items []string, multi bool) []string {
	cmdArgs := []string{}
	if multi {
		cmdArgs = append(cmdArgs, "--multi")
	}

	cmd := exec.Command("fzf", cmdArgs...)
	cmd.Stdin = strings.NewReader(strings.Join(items, "\n"))
	cmd.Stdout = &strings.Builder{}
	cmd.Stderr = os.Stderr

	var out strings.Builder
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Fzf selection cancelled or failed:", err)
		os.Exit(0)
	}

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	return lines
}
