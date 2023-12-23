// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/system/ls.go
// Original timestamp: 2023/11/20 14:12

package system

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
	"strings"
)

var VerboseOutput bool

// Info : show server- and client-related information
func Info() error {
	//if DaemonOnly && ClientOnly {
	//	return helpers.CustomError{"You cannot set both Daemononly (-d) and ClientOnly (-c) at the same time"}
	//}

	var cInfo types.Info
	var err error
	var apiver float32

	cli := auth.ClientConnect(true)
	if cInfo, err = cli.Info(context.Background()); err != nil {
		return err
	}

	fmt.Printf("\n%s info\n===========\n", helpers.Blue("DAEMON"))
	if apiver, err = CheckAPIversion(); err == nil {
		fmt.Printf("• API version: v%v\n", apiver)
	}
	fmt.Printf("• Containers: %v\n", cInfo.Containers)
	if VerboseOutput {
		fmt.Printf("   Running: %v\n", cInfo.ContainersRunning)
		fmt.Printf("   Stopped: %v\n", cInfo.ContainersStopped)
		fmt.Printf("   Paused: %v\n", cInfo.ContainersPaused)
	}
	fmt.Printf("• Images: %v\n", cInfo.Images)
	fmt.Printf("• Server version: %s\n", cInfo.ServerVersion)
	fmt.Printf("• Storage driver: %s\n", cInfo.Driver)
	if VerboseOutput {
		fmt.Printf("   Driver info: %s\n", cInfo.DriverStatus)
	}
	fmt.Printf("• Plugins: %s\n", cInfo.Plugins)

	fmt.Printf("• Swarm: %s\n", cInfo.Swarm.LocalNodeState)
	if cInfo.Swarm.LocalNodeState != "inactive" && VerboseOutput {
		fmt.Printf("   Node ID: %s\n", cInfo.Swarm.NodeID)
		fmt.Printf("   Node Address: %s\n", cInfo.Swarm.NodeAddr)
		fmt.Printf("   Control available: %t\n", cInfo.Swarm.ControlAvailable)
		fmt.Printf("   Error: %s\n", cInfo.Swarm.Error)
		fmt.Printf("   Remote managers: %s\n", cInfo.Swarm.RemoteManagers)
		fmt.Printf("   Nodes: %v\n", cInfo.Swarm.Nodes)
		fmt.Printf("   Managers: %v\n", cInfo.Swarm.Managers)
		fmt.Printf("   Cluster: %s\n", cInfo.Swarm.Cluster)
		fmt.Printf("   Warnings: %s\n", cInfo.Swarm.Warnings)
	}
	fmt.Printf("• Default runtime: %s\n", cInfo.DefaultRuntime)

	if VerboseOutput {
		fmt.Println("• Available runtimes:")
		// unmap the Runtimes data
		for key, _ := range cInfo.Runtimes {
			fmt.Printf("   %s\n", key)
		}
	}

	// ugly temp hacks
	c := strings.Fields(strings.TrimSpace(strings.Trim(fmt.Sprintf("%s", cInfo.ContainerdCommit), "{}")))
	if len(c) > 0 {
		fmt.Printf("• Containerd version: %s\n", c[0])
	}
	d := strings.Fields(strings.TrimSpace(strings.Trim(fmt.Sprintf("%s", cInfo.RuncCommit), "{}")))
	if len(d) > 0 {
		fmt.Printf("• RunC version: %s\n", d[0])
	}
	e := strings.Fields(strings.TrimSpace(strings.Trim(fmt.Sprintf("%s", cInfo.InitCommit), "{}")))
	if len(e) > 0 {
		fmt.Printf("• init version: %s\n", e[0])
	}

	fmt.Printf("• Cgroup driver: %s\n", cInfo.CgroupDriver)
	fmt.Printf("• Cgroup version: %s\n", cInfo.CgroupVersion)
	fmt.Println("• Host:")
	fmt.Printf("   Hostname: %s\n", cInfo.Name)
	fmt.Printf("   OS type: %s\n", cInfo.OSType)
	fmt.Printf("   Operating system: %s\n", cInfo.OperatingSystem)
	fmt.Printf("   Kernel version: %s\n", cInfo.KernelVersion)
	fmt.Printf("   Architecture: %s\n", cInfo.Architecture)
	fmt.Printf("   Number of CPUs: %v\n", cInfo.NCPU)
	fmt.Printf("   Total memory: %2.3f GB\n", (float64)(cInfo.MemTotal)/1024/1024/1024)
	fmt.Printf("   OOMkiller disabled: %t\n", cInfo.OomKillDisable)
	fmt.Printf("• Docker root directory: %s\n", cInfo.DockerRootDir)
	fmt.Printf("• Debug mode: %t\n", cInfo.Debug)
	fmt.Printf("• Experimental builder: %t\n", cInfo.ExperimentalBuild)
	fmt.Printf("• Live restore enabled: %t\n", cInfo.LiveRestoreEnabled)

	return nil
}
