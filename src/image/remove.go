// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// remove.go, jfgratton : 2023-11-18

package image

import (
	"context"
	"dtools/auth"
	"dtools/container"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strconv"
	"strings"
)

// RemoveImage: removes one or multiple docker image
// dtool rmi image1:tag1 image2:tag2, etc
func RemoveImage(args []string) {
	//var deleteResponse []types.ImageDeleteResponseItem
	nRemovedImages := 0
	cli := auth.ClientConnect(true)

	for _, imgList := range args {
		images, _ := cli.ImageList(context.Background(), types.ImageListOptions{})
		for _, image := range images {
			if image.ID[7:19] == imgList { // ID[7:19] corresponds to the "Image ID" column in dtools lsi
				nRemovedImages += remove(context.Background(), cli, imgList)
			}
			for _, tag := range image.RepoTags {
				// no use regexing the next instructions...
				if strings.HasSuffix(tag, ":latest") {
					tag = tag[:strings.Index(tag, ":latest")]
				}
				if tag == imgList {
					nRemovedImages += remove(context.Background(), cli, imgList)
				}
			}
		}
	}
	if nRemovedImages == 0 {
		fmt.Printf("Removed %s image. Did you mispell the name(s) ?\n", helpers.Red("0"))
	} else {
		fmt.Printf("Removed %s image(s).\n", helpers.Green(strconv.Itoa(nRemovedImages)))
	}
}

func remove(ctx context.Context, cli *client.Client, image string) int {
	var err error
	nRemovedImages := 0
	if !ForceRemoval && container.GetRunningContainersForImage(image) > 0 {
		fmt.Printf("Cannot remove %s : image has at least one container. Use the -f option to force removal.\n", helpers.Red(image))
		return nRemovedImages
	} else {
		_, err = cli.ImageRemove(ctx, image, types.ImageRemoveOptions{Force: ForceRemoval, PruneChildren: false})
	}
	if err != nil {
		fmt.Println(err)
	} else {
		nRemovedImages++
		fmt.Printf("Image removal of %s is successful.\n", helpers.Green(image))
	}
	return nRemovedImages
}
