// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/containerHelpers.go
// Original timestamp: 2023/11/12 21:27

package containers

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	cerr "github.com/jeanfrancoisgratton/customError"
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

// Returns all containers names in a sliced string
// NOTE: This function might become redundant with the usage of FilterContainersByStatus()
func getContainerNames() ([]string, *cerr.CustomError) {
	var containers []types.Container
	var ce *cerr.CustomError
	if containers, ce = ListContainers(false); ce != nil {
		return nil, ce
	}
	var containerNames []string

	for _, container := range containers {
		containerNames = append(containerNames, container.Names[0][1:])
	}
	return containerNames, nil
}

// Returns the number of running containers for the given image
func GetRunningContainersForImage(imageID string) (int, *cerr.CustomError) {
	var containers []types.Container
	var ce *cerr.CustomError
	numContainers := 0

	if containers, ce = ListContainers(false); ce != nil {
		return -1, ce
	}

	for _, container := range containers {
		//containerimg := getImageTag(containers.Image)
		if getImageTag(container.Image) == imageID || container.ImageID == imageID {
			numContainers++
		}
	}
	return numContainers, nil
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

// Unused so far, but maybe later: get the containers ID (the hex value) from its name
func getContainerID(containerName string) (string, *cerr.CustomError) {
	cli := auth.ClientConnect(false)

	// Inspect the containers to get its ID
	containerInfo, err := cli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		return "", &cerr.CustomError{Title: err.Error()}
	}

	return containerInfo.ID, nil
}

// FilterContainersByStatus
func FilterContainersByStatus(status string) []string {
	var containerList []types.Container
	var ce *cerr.CustomError

	if containerList, ce = ListContainers(false); ce != nil {
		ce.Error()
	}
	var filtered []string

	for _, container := range containerList {
		if container.State == status {
			filtered = append(filtered, container.Names[0][1:])
		}
	}
	return filtered
}

// mapNameToID() : fetches the containers ID from the hashed containers name
// Basically, we need this function because most dtools functions use human-readable names, while the SDK mostly uses
// hashes (IDs). We need a way to "translate" those names/IDs
func MapNameToId(cli *client.Client, containerName string) (string, *cerr.CustomError) {
	containerInfo, err := cli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		return "", &cerr.CustomError{Title: "Unable to map network name to ID", Message: err.Error()}
	}

	return containerInfo.ID, nil
}

func getComposeStackName(cli *client.Client, containerID string) (string, *cerr.CustomError) {
	isStack := false
	containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", &cerr.CustomError{Title: err.Error()}
	}
	labels := containerInfo.Config.Labels
	_, isStack = labels["com.docker.compose.project"]

	if isStack {
		return containerInfo.Config.Labels["com.docker.compose.project"], nil
	} else {
		return "", nil
	}
}
