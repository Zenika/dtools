// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/extras/catalog.go
// Original timestamp: 2023/12/04 21:12

package extras

import (
	"dtools/helpers"
	"encoding/json"
	"fmt"
)

// GetCatalog() : this is my equivalent of dockergetcatalog.sh, which lists all images hosted in a remote registry
func GetCatalog(remoteRegistry string) error {
	if err := FindRemoteRegistry(&remoteRegistry); err != nil {
		return err
	}

	jsonData, err := fetchJSON(remoteRegistry + "v2/_catalog")
	if err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(jsonData, "", " ")
	if err != nil {
		return helpers.CustomError{"Error formatting JSON: " + err.Error()}
	}

	fmt.Println(string(jsonBytes))
	return nil
}
