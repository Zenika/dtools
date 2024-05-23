// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/ls.go
// Original timestamp: 2023/11/13 22:14

package image

import (
	"context"
	"dtools/auth"
	"dtools/containers"
	"fmt"
	"github.com/docker/docker/api/types/image"
	cerr "github.com/jeanfrancoisgratton/customError"
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
func ListImages() *cerr.CustomError {
	var imageInfoSlice []imageInfoStruct
	var imageInfo imageInfoStruct

	ctx := context.Background()
	cli := auth.ClientConnect(true)

	images, err := cli.ImageList(ctx, image.ListOptions{All: ImageShowAll})
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		if errmsg == "Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?" {
			return &cerr.CustomError{Title: "Cannot connect to the daemon docker", Message: "Is the daemon running ?"}
		} else {
			return &cerr.CustomError{Title: "Panic", Message: err.Error()}
		}
	}

	// 1. Iterate through all images and fetch all of their tags
	for _, imgrng := range images {
		for _, tag := range imgrng.RepoTags {
			var ce *cerr.CustomError
			// 2. Iterate through all tags and collect the information
			imageInfo.reponame, imageInfo.tag = splitURI(tag)
			imageInfo.id = imgrng.ID[7:] // FIXME: [7:] is to get rid of "sha256:" .. we might need to get _that_ refined
			// Then we add creation time & size
			imageInfo.created = time.Unix(imgrng.Created, 0).Format("2006.01.02 15:04:05")
			imageInfo.size = imgrng.Size
			imageInfo.formattedSize = formatImageSize(imgrng.Size)
			if imageInfo.nContainers, ce = containers.GetRunningContainersForImage(tag); ce != nil {
				return ce
			}

			imageInfoSlice = append(imageInfoSlice, imageInfo)
		}
	}

	// 3. We now print the results
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Repository/imgrng name", "Image tag", "Image ID", "Creation time", "Size", "# containers"})
	for _, imgspec := range imageInfoSlice {
		// This is a design decision: I'll take only the first name in the containers slice
		t.AppendRow([]interface{}{imgspec.reponame, imgspec.tag, imgspec.id[:12], imgspec.created, imgspec.formattedSize, imgspec.nContainers})
	}
	t.SortBy([]table.SortBy{
		{Name: "Image name", Mode: table.Asc}})
	t.SetStyle(table.StyleDefault)
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[5] {
		case 0:
			return text.Colors{text.FgWhite}
		default:
			return text.Colors{text.FgHiYellow}
		}
		return nil
	})
	t.Render()

	return nil
}
