// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/logs.go
// Original timestamp: 2023/11/14 19:33

package container

import (
	"context"
	"dtools/auth"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func Log(containerName string) error {
	cli := auth.ClientConnect(true)

	logOptions := types.ContainerLogsOptions{
		ShowStdout: StdOut,
		ShowStderr: StdErr,
		Follow:     Follow,
		Tail:       "all",
	}

	logsReader, err := cli.ContainerLogs(context.Background(), containerName, logOptions)
	if err != nil {
		panic(err)
	}
	defer logsReader.Close()

	// Read and print the logs to standard output
	_, err = io.Copy(os.Stdout, logsReader)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return nil
}
