// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/container/ls.go
// Original timestamp: 2023/11/12 21:21

package container

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strings"
	"time"
)

func ListContainers(showDaemonInfo bool) []types.Container {
	clo := types.ContainerListOptions{Size: true, All: true, Latest: true}
	cli := auth.ClientConnect(showDaemonInfo)

	containers, err := cli.ContainerList(context.Background(), clo)
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		if strings.HasPrefix(errmsg, "Cannot connect to the Docker system at") {

			fmt.Printf("Unable to connect to %s. Is the Docker system running ?\n", helpers.Red(auth.ConnectURI))
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	if !showDaemonInfo {
		return containers
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Container ID", "Image", "Name", "Created", "Exposed ports", "State", "Status"})
	for _, container := range containers {
		// This is a design decision: I'll take only the first name in the container slice
		cn := container.Names[0]
		containerImage := getImageTag(container.Image)
		ports := prettifyPortsList(container.Ports)
		t.AppendRow([]interface{}{container.ID[:10], containerImage, cn[1:], time.Unix(container.Created, 0).Format("2006.01.02 15:04:05"), ports, container.State, container.Status})
	}
	t.SortBy([]table.SortBy{
		{Name: "Container name", Mode: table.Asc},
	})
	if helpers.PlainOutput {
		t.SetStyle(table.StyleDefault)
	} else {
		t.SetStyle(table.StyleBold)
	}
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[5] {
		case "running":
			//return text.Colors{text.BgBlack, text.FgHiGreen}
			return text.Colors{text.FgHiGreen}
		case "crashed":
			return text.Colors{text.BgBlack, text.FgHiRed}
		case "blocked":
		case "suspended":
		case "paused":
			return text.Colors{text.FgHiYellow}
		}
		return nil
	})
	t.Render()
	return nil
}
