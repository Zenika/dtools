// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/exec.go
// Original timestamp: 2023/12/13 23:37

package containers

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types"
	cerr "github.com/jeanfrancoisgratton/customError"
	"golang.org/x/term"
	"io"
	"os"
)

func ExecContainer(containerName string, command []string) *cerr.CustomError {
	cID := ""
	var ce *cerr.CustomError
	ctx := context.Background()
	cli := auth.ClientConnect(true)

	// Get the file descriptor for stdout, preserve the terminal's state
	fd := int(os.Stdout.Fd())
	// Get the current terminal state and disable echo
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return &cerr.CustomError{Title: "Unable to get the terminal state:", Message: err.Error()}
	}
	defer func() {
		_ = term.Restore(fd, oldState)
	}()

	// First, we need to get the containerID
	if cID, ce = MapNameToId(cli, containerName); ce != nil {
		return ce
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
		return &cerr.CustomError{Title: "Unable to create the containers:", Message: err.Error()}
	}

	execID := resp.ID
	respStart, err := cli.ContainerExecAttach(ctx, execID, types.ExecStartCheck{
		Tty: Tty,
	})
	if err != nil {
		return &cerr.CustomError{Title: "Unable to start the containers:", Message: err.Error()}
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
		return &cerr.CustomError{Title: "Unable to inspect the containers:", Message: err.Error()}
	}

	if respInspect.ExitCode != 0 {
		fmt.Printf("Command exited with non-zero status: %d\n", respInspect.ExitCode)
	}
	return nil
}
