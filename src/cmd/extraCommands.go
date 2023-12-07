// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/extraCommands.go
// Original timestamp: 2023/11/14 19:47

package cmd

import (
	"dtools/extras"
	"dtools/repo"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var getCommand = &cobra.Command{
	Use:   "get { catalog | tags }",
	Short: "Lists all images/tags in remote registry",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: dtools get { catalog | tags }")
	},
}

var catalogCommand = &cobra.Command{
	Use:   "catalog",
	Short: "Lists all images in remote registry",
	Run: func(cmd *cobra.Command, args []string) {
		remoteReg := ""
		if len(args) != 0 {
			remoteReg = args[0]
		}
		extras.GetCatalog(remoteReg)
	},
}

var tagsCommand = &cobra.Command{
	Use:     "tags",
	Example: "dtools get tags IMAGE REMOTEREGISTRY",
	Short:   "Lists all tags on a given image in the remote registry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			if len(args) == 0 {
				fmt.Println("You need to provide at least a docker image name")
				os.Exit(0)
			} else {
				extras.GetTags(args[0], "")
			}
		} else {
			extras.GetTags(args[0], args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(getCommand)
	getCommand.AddCommand(catalogCommand, tagsCommand)

	catalogCommand.PersistentFlags().BoolVarP(&repo.DefaultRegistryFlag, "defaultreg", "d", false, "Use the default registry")
	tagsCommand.PersistentFlags().BoolVarP(&repo.DefaultRegistryFlag, "defaultreg", "d", false, "Use the default registry")
}
