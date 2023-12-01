// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/ls.go
// Original timestamp: 2023/11/27 21:37

package network

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"github.com/docker/docker/api/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

func ListNetworks() error {
	cli := auth.ClientConnect(true)

	// lists all networks
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return err
	}

	// we now iterate through all containers on the docker host to find any unused networks
	usedNetworks := make(map[string]bool)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return helpers.CustomError{Message: "Unable to retrieve container list: " + err.Error()}
	}

	for _, container := range containers {
		for _, network := range container.NetworkSettings.Networks {
			usedNetworks[network.NetworkID] = true
		}
	}

	// FOR LATER USE... MAYBE
	//if UnusedOnly && UsedOnly {
	//	fmt.Printf("%s both %s and %s were invoked; ignoring both\n", helpers.Yellow("WARNING: "),
	//		helpers.Yellow("-u"), helpers.Yellow("-U"))
	//	UnusedOnly = false
	//	UsedOnly = false
	//}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Driver", "Scope"})

	for _, network := range networks {
		t.AppendRow([]interface{}{network.ID[:12], network.Name, network.Driver, network.Scope})
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
		//if usedNetworks[network.ID[:12]]
		return nil
	})

	t.Render()

	return nil
}
