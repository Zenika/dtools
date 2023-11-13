// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/stopStartRestart.go
// Original timestamp: 2023/11/12 21:30

package container

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"log"
)

// StopContainer : Stop a single or multiple container
func StopContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerStop(ctx, containername, container.StopOptions{}); err != nil {
			log.Printf("Unable to stop container %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s\n", containername, helpers.Red("STOPPED."))
	}
	return nil
}

// KillContainer : Kill a single or multiple container
func KillContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerKill(ctx, containername, "TERM"); err != nil {
			log.Printf("Unable to kill container %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s.\n", containername, helpers.Red("KILLED."))
	}
	return nil
}

// StartContainer : Start a single or multiple container
func StartContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerStart(ctx, containername, types.ContainerStartOptions{}); err != nil {
			log.Printf("Unable to start container %s: %s", containername, err)
			return err
		}

		fmt.Printf("Container %s is %s\n", containername, helpers.Green("STARTED."))
	}
	return nil
}

// RestartContainer : Restart a single or multiple container
// here's a quirk that I won't bother to deal with... : All listed container will be stopped at once, before being
// started all ot once, instead of being done once after all.
func RestartContainer(containers []string) error {
	var err error

	err = StopContainer(containers)
	err = StartContainer(containers)
	return err
}

// Killall, Stopall, Startall : wrappers around KillContainer, StopContainer, StartContainer
func Killall() error {
	KillContainer(getContainerNames())
	return nil
}

func Stopall() error {
	StopContainer(getContainerNames())
	return nil
}

func Startall() error {
	StartContainer(getContainerNames())
	return nil
}
