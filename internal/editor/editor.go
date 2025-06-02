package editor

import (
	"os"
	"os/exec"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
)

func LaunchEditor(filePath string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nvim"
	}

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	assert.ErrIsNil(err, "Failed to open $EDITOR")
}
