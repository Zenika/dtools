// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/system/systemhelpers.go
// Original timestamp: 2023/12/22 21:13

package system

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"strconv"
	"strings"
)

func CheckAPIversion() (float32, error) {
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
