// certificateManager
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/changelog.go
// Original timestamp: 2023/09/30 16:07

package helpers

import "fmt"

func ChangeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	CenterPrint("CHANGELOG")
	fmt.Println()
	CenterPrint("=========")
	fmt.Println()
	fmt.Println()

	fmt.Print(`
VERSION			DATE			COMMENT
-------			----			-------
00.72.00		2024.01.27		Added a "compose stack name" to dtools ls, GO version bump, prettify output
00.70.00		2023.12.23		dtools volume subcommands completed
00.60.00		2023.12.20		dtools exec now works
00.50.00		2023.12.14		dtools exec, mostly working
00.40.00		2023.12.07		network commands, etc
00.10.00		2023.11.18		container commands completed, img pull and rmi done
00.00.01		2023.11.11		Code reset #3
`)
}
