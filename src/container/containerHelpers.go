// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/containerHelpers.go
// Original timestamp: 2023/11/12 21:27

package container

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

var StdOut, StdErr, Follow bool
var Tty, Interactive bool
var User string

// Prettifies the ports' output
func prettifyPortsList(ports []types.Port) string {
	var portsString, sourcePort string
	for _, val := range ports {
		if val.IP == "" {
			sourcePort = ""
		} else {
			sourcePort = fmt.Sprintf("%d->", val.PublicPort)
		}
		portsString += fmt.Sprintf("%s/%s%d\n", val.Type, sourcePort, val.PrivatePort)
	}
	return portsString
}

// Returns all container names in a sliced string
// NOTE: This function might become redundant with the usage of FilterContainersByStatus()
func getContainerNames() []string {
	containers := ListContainers(false)
	var containerNames []string

	for _, container := range containers {
		containerNames = append(containerNames, container.Names[0][1:])
	}
	return containerNames
}

// Returns the number of running containers for the given image
func GetRunningContainersForImage(imageID string) int {
	numContainers := 0
	containers := ListContainers(false)

	for _, container := range containers {
		//containerimg := getImageTag(container.Image)
		if getImageTag(container.Image) == imageID || container.ImageID == imageID {
			numContainers++
		}
	}
	return numContainers
}

// Standardizes the image:tag format (ie: add :latest to name if it's omitted)
// Possible values for name:
// 1.registry/image
// 2.registry/image:tag
// 3.registry:port/image
// 4.registry:port/image:tag
// 5.image
// 6.image:tag
func getImageTag(name string) string {
	slashIndex := strings.Index(name, "/")
	columnIndex := strings.LastIndex(name, ":")

	// Cases #2, #4 and #6
	if columnIndex > slashIndex {
		return name
	}
	return name + ":latest"
}

// Unused so far, but maybe later: get the container ID (the hex value) from its name
func getContainerID(containerName string) (string, error) {
	cli := auth.ClientConnect(false)

	// Inspect the container to get its ID
	containerInfo, err := cli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		return "", err
	}

	return containerInfo.ID, nil
}

// FilterContainersByStatus
func FilterContainersByStatus(status string) []string {
	containerList := ListContainers(false)
	var filtered []string

	for _, container := range containerList {
		if container.State == status {
			filtered = append(filtered, container.Names[0][1:])
		}
	}
	return filtered
}

// mapNameToID() : fetches the container ID from the hashed container name
// Basically, we need this function because most dtools functions use human-readable names, while the SDK mostly uses
// hashes (IDs). We need a way to "translate" those names/IDs
func MapNameToId(cli *client.Client, containerName string) (string, error) {
	containerInfo, err := cli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		return "", err
	}

	return containerInfo.ID, nil
}

func getComposeStackName(cli *client.Client, containerID string) (string, error) {
	isStack := false
	containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", err
	}
	labels := containerInfo.Config.Labels
	_, isStack = labels["com.docker.compose.project"]

	if isStack {
		return containerInfo.Config.Labels["com.docker.compose.project"], nil
	} else {
		return "", nil
	}
}
