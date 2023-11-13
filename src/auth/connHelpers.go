// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/auth/connHelpers.go
// Original timestamp: 2023/10/23 19:02

package auth

import (
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/client"
	"os"
	"strings"
)

func buildConnectURI() string {
	if strings.HasPrefix(ConnectURI, "unix://") {
		return ConnectURI
	}
	if !strings.HasPrefix(ConnectURI, "tcp://") {
		if !strings.Contains(ConnectURI, ":") {
			ConnectURI = "tcp://" + ConnectURI + ":2375"
		} else {
			ConnectURI = "tcp://" + ConnectURI
		}
	} else {
		if !strings.Contains(ConnectURI, ":") {
			ConnectURI += ":2375"
		}
	}
	return ConnectURI
}

func ClientConnect(showHostinfo bool) *client.Client {
	uri := buildConnectURI()
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker client: %s\n", err)
		os.Exit(-1)
	}
	if showHostinfo {
		ShowHost(uri, showHostinfo)
	}
	return cli
}

func ShowHost(uri string, showNow bool) string {
	//if uri == "" {
	//	uri = BuildConnectURI()
	//}
	if strings.HasPrefix(uri, "unix://") {
		uri = "localhost (unix socket)"
	}
	if showNow {
		fmt.Printf("\nDocker host is: %s.\n", helpers.White(uri))
	}
	return uri
}
