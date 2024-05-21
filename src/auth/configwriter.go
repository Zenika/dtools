// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/auth/configwriter.go
// Original timestamp: 2023/10/23 20:51

package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/registry"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
)

func writeNewConfFile(cfgfile string, authcfg registry.AuthConfig) *cerr.CustomError {
	if Credentials.Password == "" {
		Credentials.Password = hf.GetPassword(fmt.Sprintf("Please enter %s's password: ", hf.White(Credentials.ServerAddress)))
	}
	configData := map[string]map[string]map[string]string{
		"auths": {
			Credentials.ServerAddress: {
				"auth": base64.StdEncoding.EncodeToString([]byte(Credentials.Username + ":" + Credentials.Password)),
			},
		},
	}
	cfgJson, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		return &cerr.CustomError{Title: "Unable to marshal config data", Message: err.Error()}
	}
	if err := os.WriteFile(cfgfile, cfgJson, 0644); err != nil {
		return &cerr.CustomError{Title: "Unable to write config file", Message: err.Error()}
	}
	return nil
}
