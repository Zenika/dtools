// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/remove.go
// Original timestamp: 2023/11/30 00:45

package volume

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"strings"
)

func RemoveVolume(volumes []string) error {
	cli := auth.ClientConnect(true)

	// Loop tru volumes
	for _, vol := range volumes {
		if err := cli.VolumeRemove(context.Background(), vol, ForceRemoval); err != nil {
			if strings.Contains(err.Error(), "Error response from daemon: remove "+vol+": volume is in use") {
				//a := "Unable to remove " + vol + " "
				//return helpers.CustomError{fmt.Sprintf("Unable to remove %s : volume is in use. Consider using -f\n%s\n", helpers.Red(vol),
				//	helpers.Yellow("Please be aware that using -f might have unintended consequences on the container using the volume !"))}
				return helpers.CustomError{fmt.Sprintf("Unable to remove volume %s: the volume is used by a container\n", helpers.Red(vol))}

			} else {
				return helpers.CustomError{fmt.Sprintf("Error removing the volume: %s", err)}
			}
		} else {
			fmt.Printf("Removed volume %s\n", helpers.Green(vol))
		}
	}
	return nil
}
