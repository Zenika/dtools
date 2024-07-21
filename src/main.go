package main

import (
	"dtools/cmd"
	"dtools/system"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
)

func main() {
	var minimalVersion float32 = 1.43
	var err *cerr.CustomError
	var goodVer float32
	if goodVer, err = system.CheckAPIversion(); err != nil {
		fmt.Println(err.Error())
	}

	if goodVer < minimalVersion {
		fmt.Printf("Expected API version: %v, got: %v. Exiting.\n", minimalVersion, goodVer)
	}
	cmd.Execute()
}
