// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/push.go
// Original timestamp: 2023/11/18 21:56

package image

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"dtools/repo"
	"fmt"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"github.com/moby/term"
	"os"
	"runtime"
	"strings"
)

// Push() : push docker images to remote registy
func Push(images []string) error {
	var reg repo.DefaultRegistryStruct
	var err error
	imageExists := false

	cli := auth.ClientConnect(true)

	if repo.DefaultRegistryFlag {
		if reg, err = repo.ReadDefaultFile(); err != nil {
			reg = repo.DefaultRegistryStruct{}
		}
	}
	authStr := ""
	if authStr, err = auth.GetAuthString(reg.Registry); err != nil {
		return err
	}

	// loop tru all command line args
	for _, argEl := range images {
		var repository string
		if reg.Registry != "" {
			repository = reg.Registry + "/"
			repository = strings.TrimPrefix(repository, "https://")
			repository = strings.TrimPrefix(repository, "http://")
		}
		fmt.Printf("Pushing image %s%s ...\n", repository, argEl)
		argEl = fiximageTag(argEl)

		// Before even trying to push, we need to ensure that the image actually exists
		imageExists, err = ImgExists(cli, repository+argEl)
		if err != nil {
			return err
		}
		if !imageExists {
			fmt.Printf("Image %s does not exist locally.\n", hf.Red(repository+argEl))
			continue
		}
		if err := push(cli, repository, argEl, authStr); err != nil {
			return err
		}
	}
	return nil
}

// The actual push function with output display
func push(cli *client.Client, repository, imgname, authStr string) error {
	pushResponse, pusherr := cli.ImagePush(context.Background(), repository+imgname, image.PushOptions{false, authStr, nil, runtime.GOARCH})
	if pusherr != nil {
		a := pusherr.Error()
		if strings.HasPrefix(a, "invalid reference format") {
			return helpers.CustomError{Message: fmt.Sprintf("You are trying to push %s. The format is invalid.\n", hf.White(imgname))}
		}
		if strings.HasPrefix(a, "Error response from daemon: push access denied") {
			fmt.Printf("%s: either the repository %s does not exist, or login access has not been provided.\n", hf.Red("Denied"), hf.White(imgname))
			return helpers.CustomError{Message: fmt.Sprintf("%s: either the repository %s does not exist, or login access has not been provided.\n", hf.Red("Denied"), hf.White(imgname))}
		}
		if strings.HasPrefix(a, "Error response from daemon: manifest for ") {
			fmt.Printf("%s %s: manifest not found\n", hf.Red("Unable to pull"), hf.Red(imgname))
			return helpers.CustomError{Message: fmt.Sprintf("%s %s: manifest not found\n", hf.Red("Unable to push"), hf.Red(imgname))}
		}
		if strings.HasSuffix(a, "connect: connection refused") {
			return helpers.CustomError{Message: fmt.Sprintf("Connection %s at %s. Are you sure that the daemon is running ?", hf.Red("REFUSED"), hf.Blue(repository[:len(repository)-1]))}
		} else {
			panic(pusherr)
		}
	}
	defer pushResponse.Close()

	termFd, isTerm := term.GetFdInfo(os.Stdout)
	jsonmessage.DisplayJSONMessagesStream(pushResponse, os.Stdout, termFd, isTerm, nil)

	fmt.Printf("%s %s\n", hf.Green("Successfully pulled"), hf.White(repository+imgname))

	return nil
}
