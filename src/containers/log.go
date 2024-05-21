// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/log.go
// Original timestamp: 2023/11/14 19:33

package containers

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types/container"
	cerr "github.com/jeanfrancoisgratton/customError"
	"io"
	"os"
)

func Log(containerName string) *cerr.CustomError {
	cli := auth.ClientConnect(true)

	logOptions := container.LogsOptions{
		ShowStdout: StdOut,
		ShowStderr: StdErr,
		Follow:     Follow,
		Tail:       "all",
	}

	logsReader, err := cli.ContainerLogs(context.Background(), containerName, logOptions)
	if err != nil {
		return &cerr.CustomError{Title: fmt.Sprintf("Error getting %s's logs:", containerName),
			Message: err.Error()}
	}
	defer logsReader.Close()

	// Read and print the logs to standard output
	_, err = io.Copy(os.Stdout, logsReader)
	if err != nil && err != io.EOF {
		return &cerr.CustomError{Title: fmt.Sprintf("Unable to print %s's logs:", containerName), Message: err.Error()}
	}
	return nil
}
