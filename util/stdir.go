package util

import (
	"bufio"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
)

func ReadRuneTTY(x *bufio.Reader, y int) (rune, int, error) {
	if terminal.IsTerminal(y) {
		oldState, err := term.MakeRaw(y)
		if err != nil {
			panic(err)
		}
		defer term.Restore(y, oldState)
	}
	return x.ReadRune()
}
