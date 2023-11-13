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

func init() {
	rootCmd.AddCommand(lsCmd, pauseCmd, unpauseCmd, renameCmd, rmCmd)
	rootCmd.AddCommand(stopCmd, killCmd, stopallCmd, killallCmd, startCmd, startCallmd, restartCmd)

	lsCmd.PersistentFlags().BoolVarP(&helpers.PlainOutput, "plain", "P", false, "Tables are shown with less decorations")
}
