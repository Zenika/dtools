// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/volume/volumeHelpers.go
// Original timestamp: 2023/12/23 15:58

package volume

import "github.com/docker/docker/api/types/filters"

// Please note: ForceRemoval does not work !
// I stepped-by-stepped Docker's SDK, and for some reason, the code just ignores the flag
var ForceRemoval = false

var DriverName = "local"

func filterArgs(volumeName string) filters.Args {
	return filters.NewArgs(filters.KeyValuePair{
		Key:   "volume",
		Value: volumeName,
	})
}
