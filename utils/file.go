package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rishabh-j-23/ex-crl/internal/assert"
)

func LoadJSONFile[T any](path string, out *T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, out)
}

func GetFile(dir, requestName string) string {
	requestFile := filepath.Join(dir, requestName)

	_, err := os.Stat(requestFile)
	assert.ErrIsNil(err, fmt.Sprintf("%s request does not exists", requestName))

	return requestFile
}

func GetProjectConfig() string {
	projectsDir := GetProjectDir()
	return GetFile(projectsDir, ProjectConfigJson)
}

// Only returns file name with .json
func SelectFile(args []string, dir string) string {
	if len(args) == 0 {
		requestsDir := dir
		return FzfSearch(requestsDir)
	} else {
		assert.EnsureNotEmpty(map[string]string{
			"request-name": args[0],
		})
		return args[0] + ".json"
	}
}
