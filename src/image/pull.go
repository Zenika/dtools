// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/pull.go
// Original timestamp: 2023/11/13 22:35

package image

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"dtools/repo"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"
	"os"
	"runtime"
	"strings"
)

// PullImage: pulls image from repository
// dtool pull image1:tag1 image2:tag2
func PullImage(args []string) {
	var reg repo.DefaultRegistryStruct
	var err error
	bAllImages := false // temp assignment until we process []args; we might even push bAllImages as a global variable at some point
	ctx := context.Background()
	cli := auth.ClientConnect(true)

	if repo.DefaultRegistryFlag {
		if reg, err = repo.ReadDefaultFile(); err != nil {
			reg = repo.DefaultRegistryStruct{}
		}
	}
	pullOptions := types.ImagePullOptions{bAllImages, "", nil, runtime.GOARCH}

	for _, argElement := range args {
		var repo string
		if reg.Registry != "" {
			repo = reg.Registry
			if strings.HasPrefix(reg.Registry, "https://") {
				repo = strings.TrimPrefix(reg.Registry, "https://")
			}
			if strings.HasPrefix(reg.Registry, "http://") {
				repo = strings.TrimPrefix(reg.Registry, "http://")
			}
		}
		fmt.Printf("Pulling image %s...\n", argElement)
		argElement = fiximageTag(argElement)
		pullResponse, err := cli.ImagePull(ctx, repo+"/"+argElement, pullOptions)
		if err != nil {
			a := err.Error()
			if strings.HasPrefix(a, "Error response from daemon: pull access denied") {
				fmt.Printf("%s: either the repository %s does not exist, or login access has not been provided.\n", helpers.Red("Denied"), helpers.White(argElement))
				continue
			}
			if strings.HasPrefix(a, "Error response from daemon: manifest for ") {
				fmt.Printf("%s %s: manifest not found\n", helpers.Red("Unable to pull"), helpers.Red(argElement))
				continue
			} else {
				panic(err)
			}
		}
		defer pullResponse.Close()

		termFd, isTerm := term.GetFdInfo(os.Stdout)
		jsonmessage.DisplayJSONMessagesStream(pullResponse, os.Stdout, termFd, isTerm, nil)

		fmt.Printf("%s %s\n", helpers.Green("Successfully pulled"), helpers.Normal(argElement))
	}
}
