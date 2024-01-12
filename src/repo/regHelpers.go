// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/repo/regHelpers.go
// Original timestamp: 2023/11/13 22:39

package repo

import (
	"encoding/json"
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

func (payload DefaultRegistryStruct) ReadDefaultFile() error {
	//var payload = DefaultRegistryStruct{"", "", ""}

	jsonfile, err := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "dtools", "defaults.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonfile, &payload)
	if err != nil {
		return err
	}
	return nil
}
