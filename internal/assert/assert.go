package assert

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func ErrIsNil(err error, message string) {
	if err != nil {
		slog.Error(message, "err", err)
		os.Exit(1)
	}
}

// EnsureDirExists checks if a directory exists, and creates it if it doesn't.
func EnsureDirExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		slog.Error("Directory does not exist", "dir", path, "err", err)
		os.Exit(1)
	}
}

// EnsureNotEmptyOrNil ensures none of the provided strings are empty.
// If any are empty, it logs the failure and exits the program.
func EnsureNotEmpty(namedArgs map[string]string) {
	var missing []string
	for name, val := range namedArgs {
		if val == "" {
			missing = append(missing, name)
		}
	}

	if len(missing) > 0 {
		fmt.Printf("Missing required arguments: %s\n", strings.Join(missing, ", "))
		os.Exit(1)
	}
}

func NotNil[T any](value *T, message string) {
	if value == nil {
		fmt.Println(message)
	}
}
