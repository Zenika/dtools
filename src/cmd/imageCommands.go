// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/imageCommands.go
// Original timestamp: 2023/11/13 22:17

package cmd

import (
	"dtools/image"
	"dtools/repo"
	"fmt"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:     "image",
	Aliases: []string{"img"},
	Short:   "Image subcommands",
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
	Long:    `Similar to docker image, this will give you an inventory of all images on the hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		allImages := false
		if len(args) > 0 && args[0] == "all" {
			allImages = true
		}
		image.ListImages(allImages)
	},
}

var imgRmCmd = &cobra.Command{
	Use:     "rmi",
	Aliases: []string{"imagerm", "imgrm"},
	Short:   "Image remove",
	Long: `Similar to docker image rm, this will remove an image from the local inventory
FIXME FIXME FIXME : full images might not be removed; NEEDS CHECKING.`,
	Run: func(cmd *cobra.Command, args []string) {
		image.RemoveImage(args)
	},
}

var imgPullCmd = &cobra.Command{
	Use:     "pull",
	Aliases: []string{"fetch"},
	Short:   "Pulls an image from a registry",
	Long:    `Works exactly like docker pull.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.PullImage(args); err != nil {
			_ = err.Error()
		}
	},
}

var imgPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes an image to a registry",
	Long:  `Works exactly like docker push.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.Push(args); err != nil {
			_ = err.Error()
		}
	},
}

var imgTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tags an existing image with a new tag",
	Long:  `Works exactly like docker tag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.Tag(args[0], args[1]); err != nil {
			_ = err.Error()
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd, imgLsCmd, imgPullCmd, imgPushCmd, imgTagCmd, imgRmCmd)
	imageCmd.AddCommand(imgLsCmd, imgPullCmd, imgPushCmd, imgTagCmd, imgRmCmd)

	imgTagCmd.PersistentFlags().BoolVarP(&image.OverwriteTag, "overwritetag", "o", false, "If tag already exists, ")
	imgPullCmd.PersistentFlags().BoolVarP(&repo.DefaultRegistryFlag, "defaultreg", "d", false, "Use the default registry")
	imgPushCmd.PersistentFlags().BoolVarP(&repo.DefaultRegistryFlag, "defaultreg", "d", false, "Use the default registry")
}
