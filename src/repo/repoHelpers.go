// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/repo/repoHelpers.go
// Original timestamp: 2023/11/13 22:39

package repo

import (
	"encoding/json"
	cerr "github.com/jeanfrancoisgratton/customError"
	"os"
	"path/filepath"
)

type DefaultRegistryStruct struct {
	Registry string `json:"registry"`
	Username string `json:"username,omitempty"`
	Comments string `json:"comments,omitempty"`
}

var RegistryInfo = DefaultRegistryStruct{
	Registry: "https://index.docker.io/v1/",
	Username: os.Getenv("USER")}

var DefaultRegistryFlag = false

func ReadDefaultFile() (DefaultRegistryStruct, *cerr.CustomError) {
	var payload DefaultRegistryStruct
	//var payload = DefaultRegistryStruct{"", "", ""}

	jsonfile, err := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "dtools", "defaultRegistry.json"))
	if err != nil {
		return DefaultRegistryStruct{}, &cerr.CustomError{Title: "Unable to read default registry file", Message: err.Error()}
	}

	err = json.Unmarshal(jsonfile, &payload)
	if err != nil {
		return DefaultRegistryStruct{}, &cerr.CustomError{Title: "Unable to unmarshal data", Message: err.Error()}
	}
	return payload, nil
}
