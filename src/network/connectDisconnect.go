// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/connectDisconnect.go
// Original timestamp: 2023/12/04 19:04

package network

import (
	"context"
	"dtools/auth"
	"dtools/container"
	"dtools/helpers"
)

func ConnectNetwork(networkName, containerName string) error {
	var nID, cID string
	var err error

	cli := auth.ClientConnect(true)

	if nID, err = MapNameToId(cli, networkName); err != nil {
		return helpers.CustomError{"Unable to map network name to network ID: " + err.Error()}
	}
	if cID, err = container.MapNameToId(cli, containerName); err != nil {
		return helpers.CustomError{"Unable to map container name to container ID: " + err.Error()}
	}

	return cli.NetworkConnect(context.Background(), nID, cID, nil)
}

func DisconnectNetwork(networkName, containerName string) error {
	var nID, cID string
	var err error

	cli := auth.ClientConnect(true)

	if nID, err = MapNameToId(cli, networkName); err != nil {
		return helpers.CustomError{"Unable to map network name to network ID: " + err.Error()}
	}
	if cID, err = container.MapNameToId(cli, containerName); err != nil {
		return helpers.CustomError{"Unable to map container name to container ID: " + err.Error()}
	}

	return cli.NetworkDisconnect(context.Background(), nID, cID, ForceDisconnect)
}
