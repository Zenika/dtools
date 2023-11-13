// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/containerCommands.go
// Original timestamp: 2023/11/12 21:23

package cmd

import (
	"dtools/container"
	"dtools/helpers"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"lsc", "containerls"},
	Short:   "Lists all containers",
	Long:    `Equivalent to docker ps -a.`,
	Run: func(cmd *cobra.Command, args []string) {
		container.ListContainers(true)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.PersistentFlags().BoolVarP(&helpers.PlainOutput, "plain", "P", false, "Tables are shown with less decorations")
}
