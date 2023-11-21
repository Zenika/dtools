// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/ls.go
// Original timestamp: 2023/11/13 22:14

package image

import (
	"context"
	"dtools/auth"
	"dtools/container"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"
)

// CURRENT FORMAT:
// ---------------
// REPOSITORY              TAG        IMAGE ID       CREATED        SIZE
// nexus:9820/nginx        16.85.00   d15176bc14c2   28 hours ago   86MB
// nexus:9820/nginx        latest     d15176bc14c2   28 hours ago   86MB

// f*ckin' huge mess in here.... :(
func ListImages(allImg bool) {
	var imageInfoSlice []imageInfoStruct
	var imageInfo imageInfoStruct

	ctx := context.Background()
	cli := auth.ClientConnect(true)

	images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		if errmsg == "Cannot connect to the Docker system at unix:///var/run/docker.sock. Is the docker system running?" {
			fmt.Println(errmsg)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	// 1. Iterate throught all images and fetch all their tags
	for _, image := range images {
		for _, tag := range image.RepoTags {
			// 2. Iterate throught all tags and collect the information
			imageInfo.reponame, imageInfo.tag = splitURI(tag)
			imageInfo.id = image.ID[7:] // FIXME: [7:] is to get rid of "sha256:" .. we might need to get _that_ refined
			// Then we add creation time & size
			imageInfo.created = time.Unix(image.Created, 0).Format("2006.01.02 15:04:05")
			imageInfo.size = image.Size
			imageInfo.formattedSize = formatImageSize(image.Size)
			imageInfo.nContainers = container.GetRunningContainersForImage(tag)

			imageInfoSlice = append(imageInfoSlice, imageInfo)
		}
	}

	// 3. We now print the results
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Repository + image name", "Image tag", "Image ID", "Creation time", "Size", "# containers"})
	for _, imgspec := range imageInfoSlice {
		// This is a design decision: I'll take only the first name in the container slice
		t.AppendRow([]interface{}{imgspec.reponame, imgspec.tag, imgspec.id[:12], imgspec.created, imgspec.formattedSize, imgspec.nContainers})
	}
	t.SortBy([]table.SortBy{
		{Name: "Image name", Mode: table.Asc}})
	if helpers.PlainOutput {
		t.SetStyle(table.StyleDefault)
	} else {
		t.SetStyle(table.StyleBold)
	}
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}
