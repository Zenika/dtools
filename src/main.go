package main

import (
	"context"
	"dtools/auth"
	"dtools/cmd"
	"dtools/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var minimalVersion float32 = 1.43
	var err error
	var goodVer float32
	if goodVer, err = checkAPIversion(); err != nil {
		fmt.Printf("Unable to fetch API version: %s", err)
		os.Exit(0)
	}

	if goodVer < minimalVersion {
		fmt.Printf("Expected API version: %v, got: %v. Exiting.\n", minimalVersion, goodVer)
	}
	cmd.Execute()
}

// checkAPIversion() : If the installed Docker API version is below our requirements, we bail out
func checkAPIversion() (float32, error) {
	cli := auth.ClientConnect(false)

	// Get Docker server version
	version, err := cli.ServerVersion(context.Background())
	if err != nil {
		return 0.0, helpers.CustomError{Message: "Failed to get Docker server version:" + err.Error()}
	}

	// Parse the installed Docker API version to compare
	installedVersion, err := strconv.ParseFloat(strings.TrimPrefix(version.APIVersion, "v"), 64)
	if err != nil {
		return 0.0, helpers.CustomError{Message: "Failed to parse installed API version:" + err.Error()}
	}

	return (float32)(installedVersion), nil
}
