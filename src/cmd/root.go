// dtools
// src/cmd/root.go

package cmd

import (
	"dtools/auth"
	"dtools/helpers"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "dtools",
	Short:   "Docker client",
	Version: "0.00.01-0 (2023.11.11)",
	Long: `A modern-day docker client.
This tools will perform the same tasks as the official docker tool, with some extra features, especially
Where you handle remote docker repositories.`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows the Changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.ChangeLog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(clCmd)
	rootCmd.PersistentFlags().StringVarP(&auth.ConnectURI, "host", "H", "unix:///var/run/docker.sock", "Remote hosts:port to connect to")
}
