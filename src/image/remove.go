// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// remove.go, jfgratton : 2023-11-18

package image

import (
	"context"
	"dtools/auth"
	"dtools/containers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"strconv"
	"strings"
)

// RemoveImage: removes one or multiple docker image
// dtool rmi image1:tag1 image2:tag2, etc
func RemoveImage(args []string) *cerr.CustomError {
	//var deleteResponse []types.ImageDeleteResponseItem
	nRemovedImages := 0
	cli := auth.ClientConnect(true)

	for _, imgList := range args {
		images, _ := cli.ImageList(context.Background(), types.ImageListOptions{})
		for _, image := range images {
			var ce *cerr.CustomError
			if image.ID[7:19] == imgList { // ID[7:19] corresponds to the "Image ID" column in dtools lsi
				nr := 0
				if nr, ce = remove(context.Background(), cli, imgList); ce != nil {
					return ce
				} else {
					nRemovedImages += nr
				}
			}
			for _, tag := range image.RepoTags {
				// no use regexing the next instructions...
				if strings.HasSuffix(tag, ":latest") {
					tag = tag[:strings.Index(tag, ":latest")]
				}
				if tag == imgList {
					nr := 0
					if nr, ce = remove(context.Background(), cli, imgList); ce != nil {
						return ce
					} else {
						nRemovedImages += nr
					}
				}
			}
		}
	}
	if nRemovedImages == 0 {
		fmt.Printf("Removed %s image. Did you mispell the name(s) ?\n", hf.Red("0"))
	} else {
		fmt.Printf("Removed %s image(s).\n", hf.Green(strconv.Itoa(nRemovedImages)))
	}

	return nil
}

func remove(ctx context.Context, cli *client.Client, image string) (int, *cerr.CustomError) {
	var err error
	var ce *cerr.CustomError
	var nRemovedImages, nRunningContainers int

	if nRunningContainers, ce = containers.GetRunningContainersForImage(image); ce != nil {
		return 0, ce
	} else {
		if !ForceRemoval && nRunningContainers > 0 {
			return 0, &cerr.CustomError{Fatality: cerr.Warning,
				Title:   fmt.Sprintf("Cannot remove %s: there are %d running containers within", image, nRunningContainers),
				Message: "Consider using the -f option to force removal"}
		} else {
			_, err = cli.ImageRemove(ctx, image, types.ImageRemoveOptions{Force: ForceRemoval, PruneChildren: false})
		}
		if err != nil {
			return 0, &cerr.CustomError{Title: err.Error()}
		} else {
			nRemovedImages++
			fmt.Printf("Image removal of %s is successful.\n", hf.Green(image))
		}
		return nRemovedImages, nil
	}
}
