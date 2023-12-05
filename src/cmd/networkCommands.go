// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/networkCommands.go
// Original timestamp: 2023/11/30 00:05

package cmd

import (
	"dtools/network"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

var networkCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a network",
	Long:    `Similar to docker network ls, this will give you an inventory of all networks on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		network.CreateNetwork(args)
	},
}

var networkRemoveCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "del"},
	Short:   "Delete a network",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide at least one network name")
			os.Exit(0)
		}
		network.RemoveNetwork(args)
	},
}

var networkConnectCmd = &cobra.Command{
	Use:     "connect",
	Example: "connect NETWORK CONTAINER",
	Short:   "Connects a container to the network",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You must provide the network name, then the container name")
			os.Exit(0)
		}
		if err := network.ConnectNetwork(args[0], args[1]); err != nil {
			fmt.Printf("Error connecting %s to %s: %s\n", args[0], args[1], err)
		}
	},
}

var networkDisConnectCmd = &cobra.Command{
	Use:     "disconnect",
	Example: "disconnect NETWORK CONTAINER",
	Short:   "Disconnects a container to the network",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You must provide the network name, then the container name")
			os.Exit(0)
		}
		if err := network.DisconnectNetwork(args[0], args[1]); err != nil {
			fmt.Printf("Error disconnecting %s to %s: %s\n", args[0], args[1], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(networkLsCmd, networkCreateCmd, networkRemoveCmd, networkConnectCmd, networkDisConnectCmd)

	networkCreateCmd.PersistentFlags().BoolVarP(&network.Attachable, "attachable", "", false, "Enable manual container attachment")
	networkCreateCmd.PersistentFlags().StringSliceVarP(&network.AuxAddr, "aux-address", "", []string{}, "Auxiliary IPv[46] addresses used by the Network driver")
	networkDisConnectCmd.PersistentFlags().BoolVarP(&network.ForceDisconnect, "force", "f", false, "Force network disconnection even if the container uses the network")

	//networkLsCmd.PersistentFlags().BoolVarP(&network.UsedOnly, "unused", "U", false, "Unused networks only")
}
