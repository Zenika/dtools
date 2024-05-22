// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/misc.go
// Original timestamp: 2023/11/12 21:16

package helpers

const (
	terminalEscape = "\x1b"
)

// CustomError implements the error interface
type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}
