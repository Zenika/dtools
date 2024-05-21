// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/containerCommands.go
// Original timestamp: 2023/11/12 21:23

package cmd

import (
	"dtools/containers"
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
		if _, ce := containers.ListContainers(true); ce != nil {
			ce.Error()
		}
	},
}

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"down", "containerdown"},
	Short:   "Cleanly stops a running containers",
	Long:    `This will attempt to gracefully shut a containers down.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.StopContainer(args)
	},
}

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kills a running containers",
	Long:  `Will SIGTERM a running containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.KillContainer(args)
	},
}

var stopallCmd = &cobra.Command{
	Use:   "stopall",
	Short: "Cleanly stops all running containers",
	Long:  `This will attempt to gracefully shut containers down.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.Stopall()
	},
}

var killallCmd = &cobra.Command{
	Use:   "killall",
	Short: "Kills all running containers",
	Long:  `Will SIGTERM all running containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.Killall()
	},
}

var startCmd = &cobra.Command{
	Use:     "start <containerName1> [containerName2] ...[containerNameX]",
	Aliases: []string{"up", "containerup"},
	Short:   "Starts one or many stopped containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers.StartContainer(args)
	},
}

var startCallmd = &cobra.Command{
	Use:   "startall",
	Short: "Starts all containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers.Startall()
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers.RestartContainer(args)
	},
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses one or many running containers(s)",
	Long:  `n/a`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.PauseContainer(args)
	},
}

var unpauseCmd = &cobra.Command{
	Use:     "unpause",
	Aliases: []string{"resume"},
	Short:   "Resumes one or many paused containers(s)",
	Long:    `This can only used with containers in a PAUSED state.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.UnpauseContainer(args)
	},
}

var renameCmd = &cobra.Command{
	Use: "rename",
	//Aliases: []string{"execute"},
	Short: "Renames a containers",
	//Long:    `Executes a command on the named containers. Options are mostly the same as docker exec`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.RenameContainer(args[0], args[1])
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes one or many containers",
	Long:  `This will remove stopped containers(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.RemoveContainer(args)
	},
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspects a containers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must include a containers name")
			os.Exit(0)
		}
		if err := containers.Inspect(args[0]); err != nil {
			err.Error()
		}
	},
}

var logCmd = &cobra.Command{
	Use:     "log",
	Aliases: []string{"logs"},
	Short:   "Shows the containers's logs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide a containers name")
			os.Exit(-1)
		}
		if err := containers.Log(args[0]); err != nil {
			_ = err.Error()
		}
	},
}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Starts (runs) a containers",
	Example: "see dtools run -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("WIP")
			os.Exit(-1)
		}
		if err := containers.RunContainer(args); err != nil {
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
			fmt.Println("You need to provide a containers name, and then a command, with its parameters (if needed)")
			os.Exit(-1)
		}
		//if err := containers.ExecContainer(args[0], args[1:]); err != nil {
		if err := containers.ExecContainer(args[0], args[1:]); err != nil {
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
		if err := containers.DiffContainer(args); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd, pauseCmd, unpauseCmd, renameCmd, rmCmd, inspectCmd, logCmd, runCmd, execCmd)
	rootCmd.AddCommand(stopCmd, killCmd, stopallCmd, killallCmd, startCmd, startCallmd, restartCmd)

	logCmd.PersistentFlags().BoolVarP(&containers.StdOut, "stdout", "o", true, "Shows stdout")
	logCmd.PersistentFlags().BoolVarP(&containers.StdErr, "stderr", "e", true, "Shows stderr")
	logCmd.PersistentFlags().BoolVarP(&containers.Follow, "follow", "f", false, "Follows (like tail -f)")

	execCmd.Flags().BoolVarP(&containers.Tty, "tty", "t", false, "Allocate a pseudo-TTY")
	execCmd.Flags().BoolVarP(&containers.Interactive, "interactive", "i", false, "Keep STDIN open even if not attached")
	execCmd.Flags().StringVarP(&containers.User, "user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
}
