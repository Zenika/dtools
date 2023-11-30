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

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return err
	}

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
	t.Render()

	return nil
}
