// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/create.go
// Original timestamp: 2023/11/30 00:44

package volume

import (
	"context"
	"dtools/auth"
	"github.com/docker/docker/api/types/volume"
	cerr "github.com/jeanfrancoisgratton/customError"
)

func CreateVolume(volumes []string) *cerr.CustomError {
	var err error
	cli := auth.ClientConnect(true)

	for _, vol := range volumes {
		createOps := volume.CreateOptions{Driver: DriverName, Name: vol}
		_, err = cli.VolumeCreate(context.Background(), createOps)
		if err != nil {
			return &cerr.CustomError{Title: "Unable to create volume", Message: err.Error()}
		}
	}
	return nil
}
