// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/auth/connHelpers.go
// Original timestamp: 2023/10/23 19:02

package auth

import (
	"dtools/helpers"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/client"
	"os"
	"path/filepath"
	"strings"
)

func buildConnectURI() string {
	if strings.HasPrefix(ConnectURI, "unix://") {
		return ConnectURI
	}
	if !strings.HasPrefix(ConnectURI, "tcp://") {
		if !strings.Contains(ConnectURI, ":") {
			ConnectURI = "tcp://" + ConnectURI + ":2375"
		} else {
			ConnectURI = "tcp://" + ConnectURI
		}
	} else {
		if !strings.Contains(ConnectURI, ":") {
			ConnectURI += ":2375"
		}
	}
	return ConnectURI
}

func ClientConnect(showHostinfo bool) *client.Client {
	uri := buildConnectURI()
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker client: %s\n", err)
		os.Exit(-1)
	}
	if showHostinfo {
		ShowHost(uri, showHostinfo)
	}
	return cli
}

func ShowHost(uri string, showNow bool) string {
	//if uri == "" {
	//	uri = BuildConnectURI()
	//}
	if strings.HasPrefix(uri, "unix://") {
		uri = "localhost (unix socket)"
	}
	if showNow {
		fmt.Printf("\nDocker host is: %s.\n", helpers.White(uri))
	}
	return uri
}

func GetAuthString(remoteReg string) (string, error) {
	var cfgData map[string]interface{}
	authString := ""

	cfgFilepath := filepath.Join(os.Getenv("HOME"), ".docker", "config.json")

	// Read the config file
	cfgFile, err := os.ReadFile(cfgFilepath)
	if err != nil {
		return "", helpers.CustomError{Message: "Unable to load the auth file: " + err.Error()}
	}

	// Parse the config file
	if err := json.Unmarshal(cfgFile, &cfgData); err != nil {
		return "", helpers.CustomError{Message: "Unable to parse the auth file: " + err.Error()}
	}
	if auths, ok := cfgData["auths"].(map[string]interface{}); ok {
		if remoteRegAuth, ok := auths[remoteReg].(map[string]interface{}); ok {
			if authValue, ok := remoteRegAuth["auth"].(string); ok {
				authString = authValue
			} else {
				return "", helpers.CustomError{Message: fmt.Sprintf("Auth value for %s is not a string.\n", helpers.Red(remoteReg))}
			}
		} else {
			return "", helpers.CustomError{Message: fmt.Sprintf("%s section not found in config file\n", helpers.Red(remoteReg))}
		}
	} else {
		return "", helpers.CustomError{Message: "No auths section found in config file."}
	}

	return authString, nil
}
