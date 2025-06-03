package assert

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ErrIsNil(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// EnsureDirExists checks if a directory exists, and creates it if it doesn't.
func EnsureDirExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Directory does not exists; dir=%s, err=%v\n", path, err)
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
