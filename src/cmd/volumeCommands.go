// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/networkCommands.go
// Original timestamp: 2023/11/30 00:05

package cmd

import (
	"dtools/helpers"
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
		if err := volume.ListVolumes(); err != nil {
			fmt.Printf("%s\n", err)
		}
	},
}

var volumeCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a network",
	Long:    `Similar to docker network ls, this will give you an inventory of all networks on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide at least one volume name")
			os.Exit(0)
		}
		if err := volume.CreateVolume(args); err != nil {
			fmt.Printf("%s\n", err)
		}
	},
}

var volumeRemoveCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "del"},
	Short:   "Delete one or many volume(s)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide at least one volume name")
			os.Exit(0)
		}
		if err := volume.RemoveVolume(args); err != nil {
			fmt.Printf("%s\n", err)
		}
	},
}

var volumeDriverLsCmd = &cobra.Command{
	Use:     "driverlist",
	Aliases: []string{"drivers", "driverls"},
	Short:   "List all volume drivers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(helpers.Yellow("Unimplemented for now (slated for version 1.00.00+"))
		//if err := volume.VolumeDriverList(); err != nil {
		//	fmt.Printf("%s\n", err)
		//}
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeLsCmd, volumeCreateCmd, volumeRemoveCmd, volumeDriverLsCmd)

	volumeRemoveCmd.PersistentFlags().BoolVarP(&volume.ForceRemoval, "force", "f", false, "Force removal of volume, even if in use by a container")
	volumeCreateCmd.PersistentFlags().StringVarP(&volume.Driver, "driver", "d", "local", "Volume driver")
	//networkCreateCmd.PersistentFlags().BoolVarP(&network.Attachable, "attachable", "", false, "Enable manual container attachment")
	//networkCreateCmd.PersistentFlags().StringSliceVarP(&network.AuxAddr, "aux-address", "", []string{}, "Auxiliary IPv[46] addresses used by the Network driver")
	//networkDisConnectCmd.PersistentFlags().BoolVarP(&network.ForceDisconnect, "force", "f", false, "Force network disconnection even if the container uses the network")

	//networkLsCmd.PersistentFlags().BoolVarP(&network.UsedOnly, "unused", "U", false, "Unused networks only")
}
