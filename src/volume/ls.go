// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/ls.go
// Original timestamp: 2023/11/27 21:37

package volume

import (
	"dtools/auth"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/volume"
	cerr "github.com/jeanfrancoisgratton/customError"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	//"github.com/docker/docker/api/types"
	"context"
)

type volumeInfoStruct struct {
	Driver, Name, UsedBy string
}

func ListVolumes() *cerr.CustomError {
	var volInfo volumeInfoStruct
	var volInfoSlice []volumeInfoStruct

	cli := auth.ClientConnect(true)

	// List volumes
	volumes, err := cli.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		return &cerr.CustomError{Title: "Error getting volume list: %s", Message: err.Error()}
	}

	for _, volume := range volumes.Volumes {
		volInfo = volumeInfoStruct{Name: volume.Name, Driver: volume.Driver}

		// List containers using the current volume
		args := filterArgs(volume.Name)
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{Filters: args, All: true})
		if err != nil {
			return &cerr.CustomError{Title: "Unable to fetch container list: %s", Message: err.Error()}
		}

		//fmt.Println("Containers using this volume:")
		for _, container := range containers {
			volInfo.UsedBy = container.Names[0]
		}
		volInfoSlice = append(volInfoSlice, volInfo)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Volume name", "Driver", "Used by container"})
	for _, vinfo := range volInfoSlice {
		t.AppendRow([]interface{}{vinfo.Name, vinfo.Driver, vinfo.UsedBy})
	}
	t.SortBy([]table.SortBy{
		{Name: "Volume name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)

	t.Style().Format.Header = text.FormatDefault
	t.Render()
	return nil
}
