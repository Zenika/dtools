// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/auth/connection.go
// Original timestamp: 2023/10/23 18:45

package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"

	//"github.com/docker/docker/api/types"
	"context"
	"github.com/docker/docker/api/types/registry"
	"os"
	"path/filepath"
)

var ConnectURI = ""

var Credentials registry.AuthConfig

func Login(args []string) *cerr.CustomError {
	var addr string
	var configMap = make(map[string]interface{})
	ctx := context.Background()
	cli := ClientConnect(false)

	// If the server address was provided, we add it in Credentials
	if len(args) > 0 {
		Credentials.ServerAddress = args[len(args)-1]
	}

	// Encode the auth string to base64
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(Credentials.Username + ":" + Credentials.Password))

	// Set the authentication configuration for the Docker client
	authConfig := registry.AuthConfig{
		Username: Credentials.Username,
		Password: Credentials.Password,
	}
	_, err := cli.RegistryLogin(ctx, authConfig)
	if err != nil {
		return &cerr.CustomError{Title: "Unable to authenticate:", Message: err.Error()}
	}

	// Get the path to the Docker config file
	configFile := filepath.Join(os.Getenv("HOME"), ".docker", "config.json")

	// Read the current config file
	configData, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return writeNewConfFile(configFile, authConfig)
		} else {
			return &cerr.CustomError{Message: err.Error()}
		}
	}

	// Unmarshal the JSON data into a map
	//configMap := make(map[string]interface{})
	err = json.Unmarshal(configData, &configMap)
	if err != nil {
		return &cerr.CustomError{Message: err.Error()}
	}

	// Update the auth info for the given registry
	addr = Credentials.ServerAddress
	configMap["auths"].(map[string]interface{})[addr] = map[string]string{"auth": encodedAuth}

	// Marshal the updated data back into JSON format
	updatedConfigData, err := json.MarshalIndent(configMap, "", "  ")
	if err != nil {
		return &cerr.CustomError{Message: err.Error()}
	}

	// Write the updated config data back to the config file
	err = os.WriteFile(configFile, updatedConfigData, 0600)
	if err != nil {
		return &cerr.CustomError{Message: err.Error()}
	}
	fmt.Printf("Logged in and authentication information saved to %s\n", filepath.Join(os.Getenv("HOME"), ".docker", "config.json"))
	return nil
}
