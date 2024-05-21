// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/killStopStartRestart.go
// Original timestamp: 2023/11/12 21:30

package containers

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types/container"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"log"
)

// StopContainer : Stop a single or multiple containers
func StopContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerStop(ctx, containername, container.StopOptions{}); err != nil {
			log.Printf("Unable to stop containers %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s\n", containername, hf.Red("STOPPED."))
	}
	return nil
}

// KillContainer : Kill a single or multiple containers
func KillContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerKill(ctx, containername, "TERM"); err != nil {
			log.Printf("Unable to kill containers %s: %s", containername, err)
			return err
		}
		fmt.Printf("Container %s is %s.\n", containername, hf.Red("KILLED."))
	}
	return nil
}

// StartContainer : Start a single or multiple containers
func StartContainer(containers []string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerStart(ctx, containername, container.StartOptions{}); err != nil {
			log.Printf("Unable to start containers %s: %s", containername, err)
			return err
		}

		fmt.Printf("Container %s is %s\n", containername, hf.Green("STARTED."))
	}
	return nil
}

// RestartContainer : Restart a single or multiple containers
// here's a quirk that I won't bother to deal with... : All listed containers will be stopped at once, before being
// started all ot once, instead of being done once after all.
func RestartContainer(containers []string) error {
	var err error

	err = StopContainer(containers)
	err = StartContainer(containers)
	return err
}

// Killall, Stopall, Startall : wrappers around KillContainer, StopContainer, StartContainer
func Killall() error {
	KillContainer(FilterContainersByStatus("running"))
	return nil
}

func Stopall() error {
	StopContainer(FilterContainersByStatus("running"))
	return nil
}

func Startall() error {
	StartContainer(FilterContainersByStatus("exited"))
	return nil
}
