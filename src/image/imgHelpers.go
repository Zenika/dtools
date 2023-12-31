// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/imgHelpers.go
// Original timestamp: 2023/11/13 22:16

package image

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"math"
	"strings"
)

// REPOSITORY              TAG              IMAGE ID       CREATED      SIZE
// rpmbuilder              latest           d7f4c25238e4   3 days ago   414MB
// nexus:9820/rpmbuilder   10.00.00-arm64   d7f4c25238e4   3 days ago   414MB
// rocky                   test             662704dd4eee   3 days ago   301MB

type imageInfoStruct struct {
	id, reponame, tag, created, formattedSize string
	size                                      int64
	nContainers                               int
}

var ForceRemoval = false
var OutputTarball string
var QuietBuild = false
var OverwriteTag = false

func splitURI(imagetag string) (string, string) {
	tag := "latest"
	repo := ""

	slashIndex := strings.Index(imagetag, "/")
	columnIndex := strings.LastIndex(imagetag, ":")

	// This means: no remote registry
	if slashIndex == -1 {
		return imagetag[:columnIndex], imagetag[columnIndex+1:]
	} else {
		repo = imagetag[:strings.LastIndex(imagetag, ":")]
		tag = imagetag[columnIndex+1:]
	}
	return repo, tag
}

// formatImageSize() : just so the image size shows MB or GB when needed
func formatImageSize(sz int64) string {
	numSize := (float32)(sz) / 1000.0 / 1000.0 // this will give us the size in MB
	if (int)(math.Log10(float64(numSize))) > 2 {
		return fmt.Sprintf("%.3f GB", numSize/1000.0)
	} else {
		return fmt.Sprintf("%.3f MB", numSize)
	}
}

func fiximageTag(imagetag string) string {
	slashIndex := strings.Index(imagetag, "/")
	columnIndex := strings.LastIndex(imagetag, ":")

	if columnIndex == -1 || columnIndex < slashIndex {
		return imagetag + ":latest"
	} else {
		return imagetag
	}
}

// ImgExists() : checks if a given image exists locally.
// Mostly needed by push
func ImgExists(cli *client.Client, imageName string) (bool, error) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}

	for _, image := range images {
		for _, tag := range image.RepoTags {
			if strings.Contains(tag, imageName) {
				return true, nil
			}
		}
	}
	return false, nil
}

// FIXME: might need some firming up, here....
// TagExists() : checks if the given tag already exists
func TagExists(cli *client.Client, newTag string) (bool, error) {
	_, _, err := cli.ImageInspectWithRaw(context.Background(), newTag)
	if err == nil {
		return true, nil
	} else if !client.IsErrNotFound(err) {
		return false, err
	}

	return false, nil
}
