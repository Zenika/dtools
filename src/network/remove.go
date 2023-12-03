// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/remove.go
// Original timestamp: 2023/11/30 00:45

package network

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
)

// RemoveNetwork() :
// Removes one or more network from docker host
func RemoveNetwork(args []string) error {
	cli := auth.ClientConnect(true)

	for _, arg := range args {
		nID, err := mapNameToId(cli, arg)
		if err != nil {
			return err
		}
		if err = cli.NetworkRemove(context.Background(), nID); err != nil {
			return err
		}
		fmt.Printf("%s %s\n", helpers.Green("Succesfully removed"), arg)
	}
	return nil
}
