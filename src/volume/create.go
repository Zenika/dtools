// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/create.go
// Original timestamp: 2023/11/30 00:44

package volume

import (
	"context"
	"dtools/auth"
	"github.com/docker/docker/api/types/volume"
)

func CreateVolume(volumes []string) error {
	var err error
	cli := auth.ClientConnect(true)

	for _, vol := range volumes {
		createOps := volume.CreateOptions{Driver: DriverName, Name: vol}
		_, err = cli.VolumeCreate(context.Background(), createOps)
		if err != nil {
			return err
		}
	}
	return nil
}
