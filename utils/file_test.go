package utils_test

import (
	"github.com/rishabh-j-23/ex-crl/utils"

	"testing"
)

func TestSaveAndLoadDataToFile(t *testing.T) {
	tmp := t.TempDir()
	path := tmp + "/foo.json"
	type foo struct{ X int }
	data := foo{X: 42}
	utils.SaveDataToFile(path, data)
	var out foo
	err := utils.LoadJSONFile(path, &out)
	if err != nil || out.X != 42 {
		t.Errorf("Save/LoadDataToFile failed: %v, out=%+v", err, out)
	}
}
