// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/systemCommands.go
// Original timestamp: 2023/11/20 19:21

package cmd

import (
	"dtools/system"
	"fmt"
	"github.com/spf13/cobra"
)

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "System subcommands",
	//	Long: `You need to provide the subcommands: ls, pull, push, build, rm, load, save.
	//Most commands have an alternate form if the verb image is not provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide one of the following subcommands: info, [...]")
	},
}

// Shows server and client info
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows server and client info",
	Run: func(cmd *cobra.Command, args []string) {
		if err := system.Info(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().BoolVarP(&system.VerboseOutput, "verbose", "v", false, "Verbose output")
}
