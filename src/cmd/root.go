// dtools
// src/cmd/root.go

package cmd

import (
	"dtools/auth"
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "dtools",
	Short:   "Docker client",
	Version: hf.White(fmt.Sprintf("0.80.00-0-%s (2024.05.22)", runtime.GOARCH)),
	Long: `A modern-day docker client.
This tools will perform the same tasks as the official docker tool, with some extra features, especially
Where you handle remote docker repositories.`,
}

// Shows changelog
var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows the Changelog",
	Run: func(cmd *cobra.Command, args []string) {
		changeLog()
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
	rootCmd.PersistentFlags().StringVarP(&auth.ConnectURI, "host", "H", "unix:///var/run/docker.sock", "Remote host:port to connect to")

}

func changeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Println("CHANGELOG")
	fmt.Println()
	fmt.Println("=========")
	fmt.Println()
	fmt.Println()

	fmt.Print(`
VERSION			DATE			COMMENT
-------			----			-------
00.80.00		2024.05.22		Moved all functions from the helpers package to my github helperFunctions package
00.75.00		2024.05.17		New docker SDK moved many types.* data types to new data structures
00.74.02		2024.02.02		Fixed issue where the "get" subcommand was ignored
00.74.00		2024.02.01		Moved all configs in .config/JFG/dtools/ . Now supporting insecure registries
00.73.00		2024.01.27		Prettified dtools lsi as well
00.72.00		2024.01.27		Added a "compose stack name" to dtools ls, GO version bump, prettify output
00.70.00		2023.12.23		dtools volume subcommands completed
00.60.00		2023.12.20		dtools exec now works
00.50.00		2023.12.14		dtools exec, mostly working
00.40.00		2023.12.07		network commands, etc
00.10.00		2023.11.18		containers commands completed, img pull and rmi done
00.00.01		2023.11.11		Code reset #3
`)
}
