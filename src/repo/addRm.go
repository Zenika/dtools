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

func ReadDefaultFile() (DefaultRegistryStruct, error) {
	var payload = DefaultRegistryStruct{"", "", ""}
	
	jsonfile, err := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".config", "dtools", "defaults.json"))
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal(jsonfile, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func WriteDefaultFile() error {
	rcdir := filepath.Join(os.Getenv("HOME"), ".config", "dtools")
	if _, err := os.Stat(rcdir); os.IsNotExist(err) {
		os.MkdirAll(rcdir, os.ModePerm)
	}

	if _, err := os.Stat(filepath.Join(rcdir, "defaults.json")); os.IsExist(err) {
		os.Remove(filepath.Join(rcdir, "defaults.json"))
	}

	jStream, err := json.MarshalIndent(RegistryInfo, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(rcdir, "defaults.json"), jStream, 0600)
	if err != nil {
		return err
	} else {
		return nil
	}
}
