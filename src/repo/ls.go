// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/repo/ls.go
// Original timestamp: 2023/11/18 22:06

package repo

import (
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
	"path/filepath"
)

func Ls() error {
	var defaultRepo DefaultRegistryStruct
	var err error
	var jsonfile []byte

	jsonfile, err = os.ReadFile(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "dtools", "defaultRegistry.json"))
	if err != nil {
		return &cerr.CustomError{Title: "Unable to read default registry file", Message: err.Error()}
	}
	err = json.Unmarshal(jsonfile, &defaultRepo)
	if err != nil {
		return &cerr.CustomError{Title: "Unable to parse JSON: ", Message: err.Error()}
	}

	fmt.Printf("REGISTRY: %s\nUSERNAME: %s\nCOMMENTS: %s\n", hf.White(defaultRepo.Registry), hf.White(defaultRepo.Username), hf.White(defaultRepo.Comments))
	return nil
}
