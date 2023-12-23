package main

import (
	"dtools/cmd"
	"dtools/system"
	"fmt"
	"os"
)

func main() {
	var minimalVersion float32 = 1.43
	var err error
	var goodVer float32
	if goodVer, err = system.CheckAPIversion(); err != nil {
		fmt.Printf("Unable to fetch API version: %s", err)
		os.Exit(0)
	}

	if goodVer < minimalVersion {
		fmt.Printf("Expected API version: %v, got: %v. Exiting.\n", minimalVersion, goodVer)
	}
	cmd.Execute()
}
