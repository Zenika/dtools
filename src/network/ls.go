// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/ls.go
// Original timestamp: 2023/11/27 21:37

package network

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

type networkInfoStruct struct {
	ID, Name, Driver, Scope string
	Used                    bool
}

func ListNetworks() error {
	cli := auth.ClientConnect(true)

	// lists all networks
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return err
	}

	// Now mapping the network list to the above-created struct.
	// This is done to "prettify the output, below : we map the already-existing
	// Data into that struct for purely aesthetics reasons.
	networkInfoList := mapNetworks(networks, cli)

	// FOR LATER USE... MAYBE
	//if UnusedOnly && UsedOnly {
	//	fmt.Printf("%s both %s and %s were invoked; ignoring both\n", helpers.Yellow("WARNING: "),
	//		helpers.Yellow("-u"), helpers.Yellow("-U"))
	//	UnusedOnly = false
	//	UsedOnly = false
	//}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Driver", "Scope", "Used"})

	for _, network := range networkInfoList {
		t.AppendRow([]interface{}{network.ID[:12], network.Name, network.Driver, network.Scope, network.Used})
	}
	t.SortBy([]table.SortBy{
		{Name: "Name", Mode: table.Asc}})
	if helpers.PlainOutput {
		t.SetStyle(table.StyleDefault)
	} else {
		t.SetStyle(table.StyleBold)
	}
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[4] {
		case true:
			return text.Colors{text.FgHiGreen}
		}
		return nil
	})

	t.Render()

	return nil
}

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
