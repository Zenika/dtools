// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/networkCommands.go
// Original timestamp: 2023/11/30 00:05

package cmd

import (
	"dtools/network"
	"fmt"
	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Network subcommands",
	Long:    `You need to provide the subcommands: ls, create, rm.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide one of the following subcommands: ls, create, rm")
	},
}

var networkLsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"netls", "imgls", "imagelist"},
	Short:   "Network list",
	Long:    `Similar to docker network ls, this will give you an inventory of all networks on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		network.ListNetworks()
	},
}

func init() {
	rootCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(networkLsCmd)
}
