// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/rmRename.go
// Original timestamp: 2023/11/12 21:53

package container

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types/container"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"log"
	"os"
)

// RemoveContainer : removes a single or multiple container
func RemoveContainer(containers []string) error {
	ctx := context.Background()
	removeOptions := container.RemoveOptions{RemoveVolumes: true, Force: true}
	client := auth.ClientConnect(true)

	for _, containername := range containers {
		if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
			log.Printf("Unable to remove container: %s", err)
			return err
		}
		fmt.Printf("Container %s %s.\n", hf.White(containername), hf.Red("REMOVED"))
	}
	return nil //... for now
}

// RenameContainer: renames an existing container
func RenameContainer(originalName string, newName string) error {
	ctx := context.Background()
	client := auth.ClientConnect(true)

	err := client.ContainerRename(ctx, originalName, newName)

	if err == nil {
		fmt.Printf("%s %s to %s.\n", hf.Green("Successfully renamed"), hf.White(originalName), hf.White(newName))
	} else {
		fmt.Printf("%s %s %s %s: %s\n", hf.Red("Error renaming"), originalName, "to", newName, err.Error())
		os.Exit(-1)
	}
	return err
}
