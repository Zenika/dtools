// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/repo/info.go
// Original timestamp: 2023/11/18 22:06

package repo

import (
	"dtools/helpers"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Info() error {
	var defaultRepo DefaultRegistryStruct
	var err error
	var jsonfile []byte

	jsonfile, err = os.ReadFile(filepath.Join(os.Getenv("HOME"), ".config", "dtools", "defaults.json"))
	if err != nil {
		return helpers.CustomError{Message: "Unable to read defaults file"}
	}
	err = json.Unmarshal(jsonfile, &defaultRepo)
	if err != nil {
		return helpers.CustomError{Message: "Unable to parse JSON: " + err.Error()}
	}

	fmt.Printf("REGISTRY: %s\nUSENAME: %s\nCOMMENTS: %s\n", helpers.White(defaultRepo.Registry), helpers.White(defaultRepo.Username), helpers.White(defaultRepo.Comments))
	return nil
}
