package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func FzfSearch(dir string) string {

	cmd := exec.Command("fzf", "--preview", "bat --style=numbers --color=always --paging=never {}", "--style", "full")
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to select request:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}
