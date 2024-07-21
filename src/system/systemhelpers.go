// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/system/systemhelpers.go
// Original timestamp: 2023/12/22 21:13

package system

import (
	"context"
	"dtools/auth"
	cerr "github.com/jeanfrancoisgratton/customError"
	"strconv"
	"strings"
)

func CheckAPIversion() (float32, *cerr.CustomError) {
	cli := auth.ClientConnect(false)

	// Get Docker server version
	version, err := cli.ServerVersion(context.Background())
	if err != nil {
		return 0.0, &cerr.CustomError{Title: "Failed to get Docker server version:", Message: err.Error()}
	}

	// Parse the installed Docker API version to compare
	installedVersion, err := strconv.ParseFloat(strings.TrimPrefix(version.APIVersion, "v"), 64)
	if err != nil {
		return 0.0, &cerr.CustomError{Fatality: cerr.Continuable,
			Title:   "Failed to parse installed API version:",
			Message: err.Error()}
	}

	return (float32)(installedVersion), nil
}
