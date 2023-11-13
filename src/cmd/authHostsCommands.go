// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/authhostsCommands.go
// Original timestamp: 2023/10/26 21:04

package cmd

import (
	"dtools/auth"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var authCmd = &cobra.Command{
	Use:   "auth { login | logout }",
	Short: "Provides authentication services to remote registries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Usage: %s { login | logout }\n", os.Args[0])
	},
}

var loginCmd = &cobra.Command{
	Use:   "login [server]",
	Short: "Connects to remote server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := auth.Login(args); err != nil {
			fmt.Println(err)
		}
	},
}
