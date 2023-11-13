// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/rmRename.go
// Original timestamp: 2023/11/12 21:53

package container

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"log"
	"os"
)

// RemoveContainer : removes a single or multiple container
func RemoveContainer(containers []string) error {
	ctx := context.Background()
	removeOptions := types.ContainerRemoveOptions{RemoveVolumes: true, Force: true}
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
			log.Printf("Unable to remove container: %s", err)
			return err
		}
		fmt.Printf("Container %s %s.\n", helpers.White(containername), helpers.Red("REMOVED"))
	}
	return nil //... for now
}

// RenameContainer: renames an existing container
func RenameContainer(originalName string, newName string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	err := client.ContainerRename(ctx, originalName, newName)

	if err == nil {
		fmt.Printf("%s %s to %s.\n", helpers.Green("Successfully renamed"), helpers.White(originalName), helpers.White(newName))
	} else {
		fmt.Printf("%s %s %s %s: %s\n", helpers.Red("Error renaming"), originalName, "to", newName, err.Error())
		os.Exit(-1)
	}
	return err
}
