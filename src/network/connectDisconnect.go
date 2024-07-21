// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/connectDisconnect.go
// Original timestamp: 2023/12/04 19:04

package network

import (
	"context"
	"dtools/auth"
	"dtools/containers"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
)

func ConnectNetwork(networkName, containerName string) *cerr.CustomError {
	var nID, cID string
	//var err error
	var ce *cerr.CustomError

	cli := auth.ClientConnect(true)

	if nID, ce = MapNameToId(cli, networkName); ce != nil {
		return ce
	}
	if cID, ce = containers.MapNameToId(cli, containerName); ce != nil {
		return ce
	}

	if err := cli.NetworkConnect(context.Background(), nID, cID, nil); err != nil {
		return &cerr.CustomError{Title: fmt.Sprintf("Unable to connect %s to network %s", containerName, networkName),
			Message: err.Error()}
	}

	return nil
}

func DisconnectNetwork(networkName, containerName string) *cerr.CustomError {
	var nID, cID string
	var err *cerr.CustomError

	cli := auth.ClientConnect(true)

	if nID, err = MapNameToId(cli, networkName); err != nil {
		return err
	}
	if cID, err = containers.MapNameToId(cli, containerName); err != nil {
		return err
	}

	if nE := cli.NetworkDisconnect(context.Background(), nID, cID, ForceDisconnect); nE != nil {
		return &cerr.CustomError{}
	}

	return nil
}
