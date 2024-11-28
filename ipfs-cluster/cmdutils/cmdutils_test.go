package cmdutils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfigHelper(t *testing.T) {
	wd, _ := os.Getwd()
	join := filepath.Join(wd, "testdata")
	configPath := filepath.Join(join, "service.json")
	identityPath := filepath.Join(join, "identity.json")
	helper := NewConfigHelper(configPath, identityPath, "crdt", "pebble")
	defer helper.Manager().Shutdown()

	err := helper.Manager().Default()
	if err != nil {
		t.Fatal(err)
	}

	json, err := helper.Manager().ToDisplayJSON()
	if err != nil {
		t.Fatal(err)
	}
	err = helper.SaveConfigToDisk()
	if err != nil {
		t.Fatal(err)
	}
	// err = helper.SaveIdentityToDisk()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	fmt.Println(string(json))
}
