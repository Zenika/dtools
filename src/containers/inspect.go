// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/inspect.go
// Original timestamp: 2023/11/14 19:21

package containers

import (
	"context"
	"dtools/auth"
	"encoding/json"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
)

func Inspect(containerName string) *cerr.CustomError {
	cli := auth.ClientConnect(false)

	containerInfo, err := cli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		return &cerr.CustomError{Title: "Unable to inspect containers: ", Message: err.Error()}
	}
	jsonData, err := json.MarshalIndent(containerInfo, "", "    ")
	if err != nil {
		return &cerr.CustomError{Title: "Unable to marshall data in JSON format: ", Message: err.Error()}
	}

	fmt.Println(string(jsonData))

	cli.Close()
	return nil
}
