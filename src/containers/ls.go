// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/containers/ls.go
// Original timestamp: 2023/11/12 21:21

package containers

import (
	"context"
	"dtools/auth"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	cerr "github.com/jeanfrancoisgratton/customError"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strings"
	"time"
)

func ListContainers(showDaemonInfo bool) ([]types.Container, *cerr.CustomError) {
	var cli *client.Client
	clo := container.ListOptions{Size: true, All: true, Latest: true}
	if cli = auth.ClientConnect(showDaemonInfo); cli == nil {
		return nil, &cerr.CustomError{Title: "Failed to connect the Docker client"}
	}

	containers, err := cli.ContainerList(context.Background(), clo)
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		if strings.HasPrefix(errmsg, "Cannot connect to the Docker daemon at") {
			return nil, &cerr.CustomError{Title: "Cannot connect to the docker daemon:", Message: "Is the daemon runnning ?"}
		} else {
			return nil, &cerr.CustomError{Message: err.Error()}
		}
	}

	if !showDaemonInfo {
		return containers, nil
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Container ID", "Image", "Name", "Created", "Exposed ports", "State", "Status", "Compose stack"})
	for _, container := range containers {
		var composeStackName string
		var err *cerr.CustomError
		// This is a design decision: I'll take only the first name in the containers slice
		cn := container.Names[0]
		containerImage := getImageTag(container.Image)
		ports := prettifyPortsList(container.Ports)
		if composeStackName, err = getComposeStackName(cli, container.ID); err != nil {
			err.Error() // <-- same here as panic(err)
		}
		t.AppendRow([]interface{}{container.ID[:10], containerImage, cn[1:], time.Unix(container.Created, 0).Format("2006.01.02 15:04:05"), ports, container.State, container.Status, composeStackName})
	}
	t.SortBy([]table.SortBy{
		{Name: "Container name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)

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
	return containers, nil
}
