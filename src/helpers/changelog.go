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
VERSION		DATE			COMMENT
-------		----			-------
0.50.00		2023.12.14		dtools exec, mostly working
0.40.00		2023.12.07		network commands, etc
0.10.00		2023.11.18		container commands completed, img pull and rmi done
0.00.01		2023.11.11		Code reset #3
`)
}
