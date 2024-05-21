// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/pauseUnpause.go
// Original timestamp: 2023/11/12 21:48

package containers

import (
	"context"
	"dtools/auth"
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"log"
)

// PauseContainer : pauses the containers(s) given at CLI
func PauseContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerPause(ctx, containername); err != nil {
			log.Printf("Unable to stop containers %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s\n", containername, hf.Yellow("PAUSED."))
	}
	return nil
}

// UnpauseContainer : unpouses the containers(s) given at CLI
func UnpauseContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerUnpause(ctx, containername); err != nil {
			log.Printf("Unable to resume containers %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s\n", containername, hf.Green("RESUMED."))
	}
	return nil
}
