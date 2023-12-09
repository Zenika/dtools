// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/extras/extrasHelpers.go
// Original timestamp: 2023/12/04 21:27

package extras

import (
	"dtools/helpers"
	"dtools/repo"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// genericFetch() : this is the actual function that fetches the JSON payload from the URL provided
// This way, we use the same function for "dockergettags" or "dockergetcatalog"
func fetchJSON(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, helpers.CustomError{"Error getting JSON payload from url endpoint: " + err.Error()}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, helpers.CustomError{"Error reading JSON payload from url endpoint: " + err.Error()}
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// If the -d flag is used, we need to find the registered default registry to use
func FindRemoteRegistry(remoteRegistry *string) error {
	var reg repo.DefaultRegistryStruct
	var err error

	// the -d flag is used, so we need to read its config
	if repo.DefaultRegistryFlag {
		if err = reg.ReadDefaultFile(); err != nil {
			return helpers.CustomError{"Unable to read registry config file:" + err.Error()}
		}
		if reg.Registry != "" {
			*remoteRegistry = reg.Registry
		}
	}
	// The following might seem nonsensical, but is needed to ensure we have a well-formatted URL
	*remoteRegistry = strings.TrimPrefix(*remoteRegistry, "https://")
	*remoteRegistry = strings.TrimPrefix(*remoteRegistry, "http://")
	*remoteRegistry = strings.TrimSuffix(*remoteRegistry, "/")
	*remoteRegistry = "https://" + *remoteRegistry + "/"
	return nil
}
