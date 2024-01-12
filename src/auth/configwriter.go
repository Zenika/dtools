// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/auth/configwriter.go
// Original timestamp: 2023/10/23 20:51

package auth

import (
	"dtools/helpers"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

func writeNewConfFile(cfgfile string) error {
	//encodedAuth := base64.StdEncoding.EncodeToString([]byte(Credentials.Username + ":" + Credentials.Password))
	if Credentials.Password == "" {
		Credentials.Password = helpers.GetPassword(fmt.Sprintf("Please enter %s's password: ", helpers.White(Credentials.ServerAddress)))
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
		return err
	}
	if err := os.WriteFile(cfgfile, cfgJson, 0644); err != nil {
		return err
	}
	return nil
}
