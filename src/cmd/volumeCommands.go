// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/networkCommands.go
// Original timestamp: 2023/11/30 00:05

package cmd

import (
	"dtools/volume"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var volumeCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"vol"},
	Short:   "Volumes subcommands",
	Long:    `You need to provide the subcommands: ls, add, rm.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide one of the following subcommands: ls, create, rm")
	},
}

var volumeLsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"volls"},
	Short:   "Volume list",
	Long:    `Similar to docker network ls, this will give you an inventory of all networks on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		volume.ListVolumes()
	},
}

var volumeCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a network",
	Long:    `Similar to docker network ls, this will give you an inventory of all networks on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		volume.CreateVolume(args)
	},
}

var volumeRemoveCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "del"},
	Short:   "Delete a volume",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide at least one network name")
			os.Exit(0)
		}
		volume.RemoveVolume(args)
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeLsCmd, volumeCreateCmd, volumeRemoveCmd)

	//networkCreateCmd.PersistentFlags().BoolVarP(&network.Attachable, "attachable", "", false, "Enable manual container attachment")
	//networkCreateCmd.PersistentFlags().StringSliceVarP(&network.AuxAddr, "aux-address", "", []string{}, "Auxiliary IPv[46] addresses used by the Network driver")
	//networkDisConnectCmd.PersistentFlags().BoolVarP(&network.ForceDisconnect, "force", "f", false, "Force network disconnection even if the container uses the network")

	//networkLsCmd.PersistentFlags().BoolVarP(&network.UsedOnly, "unused", "U", false, "Unused networks only")
}
