// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/repo/addRm.go
// Original timestamp: 2023/11/13 22:39

package repo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteDefaultFile() error {
	rcdir := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "dtools")
	if _, err := os.Stat(rcdir); os.IsNotExist(err) {
		os.MkdirAll(rcdir, os.ModePerm)
	}

	if _, err := os.Stat(filepath.Join(rcdir, "defaultRegistry.json")); os.IsExist(err) {
		os.Remove(filepath.Join(rcdir, "defaultRegistry.json"))
	}

	jStream, err := json.MarshalIndent(RegistryInfo, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(rcdir, "defaultRegistry.json"), jStream, 0600)
	if err != nil {
		return err
	} else {
		return nil
	}
}
