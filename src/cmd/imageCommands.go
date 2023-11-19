// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/imageCommands.go
// Original timestamp: 2023/11/13 22:17

package cmd

import (
	"dtools/helpers"
	"dtools/image"
	"dtools/repo"
	"fmt"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image subcommands",
	Long: `You need to provide the subcommands: ls, pull, push, build, rm, load, save.
Most commands have an alternate form if the verb image is not provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide one of the following subcommands: lsi, pull, push, build, rmi, load, save")
	},
}

var imgLsCmd = &cobra.Command{
	Use:     "lsi",
	Aliases: []string{"imagels", "imgls", "imagelist"},
	Short:   "Image list",
	Long:    `Similar to docker image, this will give you an inventory of all image on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		allImages := false
		if len(args) > 0 && args[0] == "all" {
			allImages = true
		}
		image.ListImages(allImages)
	},
}

var imgPullCmd = &cobra.Command{
	Use:     "pull",
	Aliases: []string{"fetch", "get"},
	Short:   "Pulls an image from a registry",
	Long:    `Works exactly like docker pull.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.PullImage(args); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd, imgLsCmd, imgPullCmd)
	imageCmd.AddCommand(imgLsCmd, imgPullCmd)

	imgLsCmd.PersistentFlags().BoolVarP(&helpers.PlainOutput, "plain", "P", false, "Tables are shown with less decorations")
	imgPullCmd.PersistentFlags().BoolVarP(&repo.DefaultRegistryFlag, "defaultreg", "d", false, "Use the default registry")
}
