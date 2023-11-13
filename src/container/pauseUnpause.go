// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/pauseUnpause.go
// Original timestamp: 2023/11/12 21:48

package container

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"log"
)

// PauseContainer : pauses the container(s) given at CLI
func PauseContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerPause(ctx, containername); err != nil {
			log.Printf("Unable to stop container %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s\n", containername, helpers.Yellow("PAUSED."))
	}
	return nil
}

// UnpauseContainer : unpouses the container(s) given at CLI
func UnpauseContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerUnpause(ctx, containername); err != nil {
			log.Printf("Unable to resume container %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s\n", containername, helpers.Green("RESUMED."))
	}
	return nil
}
