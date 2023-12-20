// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/containerCommands.go
// Original timestamp: 2023/11/12 21:23

package cmd

import (
	"dtools/container"
	"dtools/helpers"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"down", "containerdown"},
	Short:   "Cleanly stops a running container",
	Long:    `This will attempt to gracefully shut a container down.`,
	Run: func(cmd *cobra.Command, args []string) {
		container.StopContainer(args)
	},
}

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kills a running container",
	Long:  `Will SIGTERM a running container.`,
	Run: func(cmd *cobra.Command, args []string) {
		container.KillContainer(args)
	},
}

var stopallCmd = &cobra.Command{
	Use:   "stopall",
	Short: "Cleanly stops all running container",
	Long:  `This will attempt to gracefully shut container down.`,
	Run: func(cmd *cobra.Command, args []string) {
		container.Stopall()
	},
}

var killallCmd = &cobra.Command{
	Use:   "killall",
	Short: "Kills all running container",
	Long:  `Will SIGTERM all running containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		container.Killall()
	},
}

var startCmd = &cobra.Command{
	Use:     "start <containerName1> [containerName2] ...[containerNameX]",
	Aliases: []string{"up", "containerup"},
	Short:   "Starts one or many stopped containers",
	Run: func(cmd *cobra.Command, args []string) {
		container.StartContainer(args)
	},
}

var startCallmd = &cobra.Command{
	Use:   "startall",
	Short: "Starts all container",
	Run: func(cmd *cobra.Command, args []string) {
		container.Startall()
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a container",
	Run: func(cmd *cobra.Command, args []string) {
		container.RestartContainer(args)
	},
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses one or many running container(s)",
	Long:  `n/a`,
	Run: func(cmd *cobra.Command, args []string) {
		container.PauseContainer(args)
	},
}

var unpauseCmd = &cobra.Command{
	Use:     "unpause",
	Aliases: []string{"resume"},
	Short:   "Resumes one or many paused container(s)",
	Long:    `This can only used with containers in a PAUSED state.`,
	Run: func(cmd *cobra.Command, args []string) {
		container.UnpauseContainer(args)
	},
}

var renameCmd = &cobra.Command{
	Use: "rename",
	//Aliases: []string{"execute"},
	Short: "Renames a container",
	//Long:    `Executes a command on the named container. Options are mostly the same as docker exec`,
	Run: func(cmd *cobra.Command, args []string) {
		container.RenameContainer(args[0], args[1])
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes one or many containers",
	Long:  `This will remove stopped container(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		container.RemoveContainer(args)
	},
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspects a container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must include a container name")
			os.Exit(0)
		}
		container.Inspect(args[0])
	},
}

var logCmd = &cobra.Command{
	Use:     "log",
	Aliases: []string{"logs"},
	Short:   "Shows the container's logs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide a container name")
			os.Exit(-1)
		}
		container.Log(args[0])
	},
}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Starts (runs) a container",
	Example: "see dtools run -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("WIP")
			os.Exit(-1)
		}
		if err := container.RunContainer(args); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
}

var execCmd = &cobra.Command{
	Use:     "exec [flags] containerID command",
	Short:   "Emulate docker exec command",
	Example: "see dtools exec -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You need to provide a container name, and then a command, with its parameters (if needed)")
			os.Exit(-1)
		}
		//if err := container.ExecContainer(args[0], args[1:]); err != nil {
		if err := container.ExecContainer(args[0], args[1:]); err != nil {
			fmt.Println("Unable to exec: ", err)
			os.Exit(-2)
		}
	},
}

var dioffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Lists diffenrces in a containers filesystem",
	//Example: "see dtools run -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("dtls diff: WIP")
			os.Exit(-1)
		}
		if err := container.DiffContainer(args); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd, pauseCmd, unpauseCmd, renameCmd, rmCmd, inspectCmd, logCmd, runCmd, execCmd)
	rootCmd.AddCommand(stopCmd, killCmd, stopallCmd, killallCmd, startCmd, startCallmd, restartCmd)

	lsCmd.PersistentFlags().BoolVarP(&helpers.PlainOutput, "plain", "P", false, "Tables are shown with less decorations")

	logCmd.PersistentFlags().BoolVarP(&container.StdOut, "stdout", "o", true, "Shows stdout")
	logCmd.PersistentFlags().BoolVarP(&container.StdErr, "stderr", "e", true, "Shows stderr")
	logCmd.PersistentFlags().BoolVarP(&container.Follow, "follow", "f", false, "Follows (like tail -f)")

	execCmd.Flags().BoolVarP(&container.Tty, "tty", "t", false, "Allocate a pseudo-TTY")
	execCmd.Flags().BoolVarP(&container.Interactive, "interactive", "i", false, "Keep STDIN open even if not attached")
	execCmd.Flags().StringVarP(&container.User, "user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
}
