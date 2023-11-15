// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/inspect.go
// Original timestamp: 2023/11/14 19:21

package container

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"encoding/json"
	"fmt"
)

func Inspect(containerName string) error {
	cli := auth.ClientConnect(false)

	containerInfo, err := cli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		return helpers.CustomError{Message: "Unable to inspect container: " + err.Error()}
	}
	jsonData, err := json.MarshalIndent(containerInfo, "", "    ")
	if err != nil {
		return helpers.CustomError{Message: "Unable to marshall data in JSON format: " + err.Error()}
	}

	fmt.Println(string(jsonData))

	cli.Close()
	return nil
}
