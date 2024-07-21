// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/extras/tags.go
// Original timestamp: 2023/12/05 19:42

package extras

import (
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
)

// GetTags() : my equivalent to dockergettags.sh, which fetches all tags of a given docker image hosted in a remote reg
func GetTags(imageName, remoteRegistry string) *cerr.CustomError {
	if err := FindRemoteRegistry(&remoteRegistry); err != nil {
		return err
	}

	jsonData, err := fetchJSON(fmt.Sprintf("%sv2/%s/tags/list", remoteRegistry, imageName))
	if err != nil {
		return err
	}

	jsonBytes, mErr := json.MarshalIndent(jsonData, "", " ")
	if mErr != nil {
		return &cerr.CustomError{Title: "Error formatting JSON: " + mErr.Error()}
	}

	fmt.Println(string(jsonBytes))
	return nil
}
