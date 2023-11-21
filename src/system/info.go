// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/system/info.go
// Original timestamp: 2023/11/20 14:12

package system

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
	"github.com/docker/docker/api/types"
)

var VerboseOutput bool

// Info : show server- and client-related information
func Info() error {
	//if DaemonOnly && ClientOnly {
	//	return helpers.CustomError{"You cannot set both Daemononly (-d) and ClientOnly (-c) at the same time"}
	//}

	var cInfo types.Info
	var err error

	cli := auth.ClientConnect(true)
	if cInfo, err = cli.Info(context.Background()); err != nil {
		return err
	}

	fmt.Printf("\n%s info\n===========\n", helpers.Blue("DAEMON"))
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

	fmt.Println("\nI'm bored.. there's a lot more to go tru (https://pkg.go.dev/github.com/docker/docker@v24.0.7+incompatible/api/types#Info)")
	fmt.Println("More later")

	return nil
}
