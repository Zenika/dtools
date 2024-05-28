// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/extras/catalog.go
// Original timestamp: 2023/12/04 21:12

package extras

import (
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// GetCatalog() : this is my equivalent of dockergetcatalog.sh, which lists all images hosted in a remote registry
func GetCatalog(remoteRegistry string) *cerr.CustomError {
	if err := FindRemoteRegistry(&remoteRegistry); err != nil {
		return &cerr.CustomError{Title: "Error reading default registry config file", Message: err.Error()}
	}

	jsonData, err := fetchJSON(remoteRegistry + "v2/_catalog")
	if err != nil {
		return err
	}

	jsonBytes, mErr := json.MarshalIndent(jsonData, "", " ")
	if mErr != nil {
		return &cerr.CustomError{Title: "Error formatting JSON: ", Message: mErr.Error()}
	}

	fmt.Println(string(jsonBytes))
	return nil
}
