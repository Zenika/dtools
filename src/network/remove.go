// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/network/remove.go
// Original timestamp: 2023/11/30 00:45

package network

import (
	"context"
	"dtools/auth"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
)

// RemoveNetwork() :
// Removes one or more network from docker host
func RemoveNetwork(args []string) *cerr.CustomError {
	var ne error
	cli := auth.ClientConnect(true)

	for _, arg := range args {
		nID, err := MapNameToId(cli, arg)
		if err != nil {
			return err
		}
		if ne = cli.NetworkRemove(context.Background(), nID); ne != nil {
			return &cerr.CustomError{Title: fmt.Sprintf("Unable to remove the %s network", arg),
				Message: ne.Error()}
		}
		fmt.Printf("%s %s\n", hf.Green("Successfully removed"), arg)
	}
	return nil
}
