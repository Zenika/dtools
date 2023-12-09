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
func PullImage(args []string) error {
	var reg repo.DefaultRegistryStruct
	var err error
	bAllImages := false // temp assignment until we process []args; we might even push bAllImages as a global variable at some point
	cli := auth.ClientConnect(true)

	if repo.DefaultRegistryFlag {
		if err = reg.ReadDefaultFile(); err != nil {
			reg = repo.DefaultRegistryStruct{}
		}
	}
	pullOptions := types.ImagePullOptions{bAllImages, "", nil, runtime.GOARCH}

	// loop tru all command line args
	for _, argElement := range args {
		var repository string
		if reg.Registry != "" {
			repository = reg.Registry + "/"
			repository = strings.TrimPrefix(repository, "https://")
			repository = strings.TrimPrefix(repository, "http://")
		}
		fmt.Printf("Pulling image %s...\n", argElement)
		argElement = fiximageTag(argElement)
		pullResponse, pullerr := cli.ImagePull(context.Background(), repository+argElement, pullOptions)
		if pullerr != nil {
			a := pullerr.Error()
			if strings.HasPrefix(a, "Error response from daemon: pull access denied") {
				fmt.Printf("%s: either the repository %s does not exist, or login access has not been provided.\n", helpers.Red("Denied"), helpers.White(argElement))
				return helpers.CustomError{Message: fmt.Sprintf("%s: either the repository %s does not exist, or login access has not been provided.\n", helpers.Red("Denied"), helpers.White(argElement))}
			}
			if strings.HasPrefix(a, "Error response from daemon: manifest for ") {
				fmt.Printf("%s %s: manifest not found\n", helpers.Red("Unable to pull"), helpers.Red(argElement))
				return helpers.CustomError{Message: fmt.Sprintf("%s %s: manifest not found\n", helpers.Red("Unable to pull"), helpers.Red(argElement))}
			}
			if strings.HasSuffix(a, "connect: connection refused") {
				return helpers.CustomError{Message: fmt.Sprintf("Connection %s at %s. Are you sure that the daemon is running ?", helpers.Red("REFUSED"), helpers.Blue(repository[:len(repository)-1]))}
			} else {
				panic(pullerr)
			}
		}
		defer pullResponse.Close()

		termFd, isTerm := term.GetFdInfo(os.Stdout)
		jsonmessage.DisplayJSONMessagesStream(pullResponse, os.Stdout, termFd, isTerm, nil)

		fmt.Printf("%s %s\n", helpers.Green("Successfully pulled"), helpers.Normal(repository+argElement))
	}
	return nil
}
