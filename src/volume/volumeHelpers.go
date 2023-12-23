// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/volume/volumeHelpers.go
// Original timestamp: 2023/12/23 15:58

package volume

import "github.com/docker/docker/api/types/filters"

func filterArgs(volumeName string) filters.Args {
	return filters.NewArgs(filters.KeyValuePair{
		Key:   "volume",
		Value: volumeName,
	})
}
