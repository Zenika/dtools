// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/exec.go
// Original timestamp: 2023/12/13 23:37

package container

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func ExecContainer(containerName string, command []string) error {
	cID := ""
	var err error
	ctx := context.Background()
	cli := auth.ClientConnect(true)

	// First, we need to get the containerID
	if cID, err = MapNameToId(cli, containerName); err != nil {
		return helpers.CustomError{fmt.Sprint("Unable to translate container name to container ID: %s", err)}
	}

	// Setup exec context
	execConfig := types.ExecConfig{
		Tty:          Tty,
		AttachStdin:  Interactive,
		AttachStdout: true,
		AttachStderr: true,
		User:         User,
		Cmd:          command,
	}

	// Create exec instance
	resp, err := cli.ContainerExecCreate(ctx, cID, execConfig)
	if err != nil {
		fmt.Printf("Error creating exec instance: %s\n", err)
		os.Exit(1)
	}

	execID := resp.ID
	respStart, err := cli.ContainerExecAttach(ctx, execID, types.ExecStartCheck{
		Tty: Tty,
	})
	if err != nil {
		fmt.Printf("Error attaching to exec instance: %s\n", err)
		os.Exit(1)
	}
	defer respStart.Close()

	// Deal with -i flag
	go func() {
		if Interactive {
			io.Copy(respStart.Conn, os.Stdin)
		}
	}()

	// Manage CTRL+D as an exit command for shells <<- looks kludged.
	go func() {
		defer respStart.Close()
		_, err := io.Copy(os.Stdout, respStart.Conn)
		if err != nil && err != io.EOF {
			fmt.Printf("Error copying data from %s: %s\n", containerName, err)
		}
	}()

	io.Copy(os.Stdout, respStart.Conn)

	// Cleanup
	respInspect, err := cli.ContainerExecInspect(ctx, execID)
	if err != nil {
		fmt.Printf("Error inspecting exec instance: %s\n", err)
		os.Exit(1)
	}

	if respInspect.ExitCode != 0 {
		fmt.Printf("Command exited with non-zero status: %d\n", respInspect.ExitCode)
	}

	return nil
}
