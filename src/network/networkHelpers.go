// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/networkHelpers.go
// Original timestamp: 2023/11/30 17:09

package network

import (
	"context"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// var UnusedOnly bool
// var UsedOnly bool
var Attachable bool
var AuxAddr []string
var ForceDisconnect = false

// mapNetworks() :
// This function is needed to determine if a network is used by any container.
// This mainly a matter of prettifying the output of `dtools network ls`
func mapNetworks(networks []types.NetworkResource, cli *client.Client) []networkInfoStruct {
	var networkInfoList []networkInfoStruct

	for _, network := range networks {
		containers, err := cli.NetworkInspect(context.Background(), network.ID, types.NetworkInspectOptions{})
		if err != nil {
			fmt.Printf("Error inspecting network %s: %s\n", helpers.Red(network.Name), err)
			continue
		}

		used := len(containers.Containers) > 0

		networkInfo := networkInfoStruct{
			ID:     network.ID,
			Name:   network.Name,
			Driver: network.Driver,
			Scope:  network.Scope,
			Used:   used,
		}
		networkInfoList = append(networkInfoList, networkInfo)
	}
	return networkInfoList
}

// mapNameToID() : fetches the network ID from the human-readable network name
// Basically, we need this function because most dtools functions use human-readable names, while the SDK mostly uses
// hashes (IDs). We need a way to "translate" those names/IDs
func MapNameToId(cli *client.Client, networkName string) (string, error) {
	networkSpecs, err := cli.NetworkInspect(context.Background(), networkName, types.NetworkInspectOptions{})
	if err != nil {
		return "", err
	}

	return networkSpecs.ID, nil
}
