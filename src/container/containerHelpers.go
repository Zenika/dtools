// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/containerHelpers.go
// Original timestamp: 2023/11/12 21:27

package container

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"strings"
)

// Prettifies the ports' output
func prettifyPortsList(ports []types.Port) string {
	var portsString, sourcePort string
	for _, val := range ports {
		if val.IP == "" {
			sourcePort = ""
		} else {
			sourcePort = fmt.Sprintf("%d->", val.PublicPort)
		}
		portsString += fmt.Sprintf("%s/%s%d  ", val.Type, sourcePort, val.PrivatePort)
	}
	return portsString
}

// Returns all container names in a sliced string
func getContainerNames() []string {
	containers := ListContainers(false)
	var containerNames []string

	for _, container := range containers {
		containerNames = append(containerNames, container.Names[0])
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
