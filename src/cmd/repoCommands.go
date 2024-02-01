// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/repoCommands.go
// Original timestamp: 2023/11/13 23:15

package cmd

import (
	"dtools/repo"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Docker registry related subcommands",
	Long:  `This is where you will find all registry-related commands, such as list tags, list image, select registry, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need one of the following subcommands: rm or add")
	},
}

var repoRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "del", "delete"},
	Short:   "Remove the default registry file",
	Long: `Removes the default registry file.
Note that this file is not necessary to get the software to work.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.Remove(filepath.Join(os.Getenv("USER"), ".config", "JFG", "dtools", "defaultRegistry.json")); err != nil {
			panic(err)
		}
	},
}

var repoAddCmd = &cobra.Command{
	Use: "add",
	//Aliases: []string{"addreg"},
	Short: "Create or overwrite the default registry file",
	Long: `This will create a default registry file in $HOME/.config/JFG/dtools/defaultRegistry.json.
The file will be overwritten if it already exists.
Note: if you do not specify the -r, -u and -c, defaults values will used.`,
	Run: func(cmd *cobra.Command, args []string) {
		if repo.RegistryInfo.Registry == "" {
			fmt.Println("The registry URL (parameter -r) must not be empty")
			os.Exit(0)
		}
		err := repo.WriteDefaultFile()
		if err != nil {
			fmt.Println(err)
		}
	},
}

var repoLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Displays info on current default repo",

	Run: func(cmd *cobra.Command, args []string) {
		repo.Ls()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoRmCmd, repoAddCmd, repoLsCmd)

	repoAddCmd.Flags().StringVarP(&repo.RegistryInfo.Registry, "registry", "r", "https://index.docker.io/v1/", "Edit the default registry")
	repoAddCmd.Flags().StringVarP(&repo.RegistryInfo.Username, "username", "u", os.Getenv("USER"), "Edit the login user to default registry")
	repoAddCmd.Flags().StringVarP(&repo.RegistryInfo.Comments, "comments", "c", "", "Add comments to default registry")
}
