// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/remove.go
// Original timestamp: 2023/11/30 00:45

package volume

import (
	"context"
	"dtools/auth"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"strings"
)

func RemoveVolume(volumes []string) *cerr.CustomError {
	cli := auth.ClientConnect(true)

	// Loop tru volumes
	for _, vol := range volumes {
		if err := cli.VolumeRemove(context.Background(), vol, ForceRemoval); err != nil {
			if strings.Contains(err.Error(), "Error response from daemon: remove "+vol+": volume is in use") {
				return &cerr.CustomError{Title: fmt.Sprintf("Unable to remove volume %s: the volume is used by a container", vol),
					Message: err.Error()}

			} else {
				return &cerr.CustomError{Title: fmt.Sprintf("Error removing volume: %s", vol),
					Message: err.Error()}
			}
		} else {
			fmt.Printf("Removed volume %s\n", hf.Green(vol))
		}
	}
	return nil
}
