package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"portal.io/dream/util"
)

func main() {
	orig := string(util.Range('𖩀', '𖩞'))
	s0 := []rune(orig)
	util.ShuffleSeed(&s0, 69)
	sa := s0[0:26]
	svb := s0[26:]
	quit := false
	r := bufio.NewReader(os.Stdin)
	alp := "abcdefghijklmnopqrstuvwxyz"
	for !quit {
		ch, _, err := util.ReadRuneTTY(r, 0)
		if err != nil {
			quit = true
		}
		if ch == 3 {
			quit = true
		}
		if !quit {
			if strings.ContainsRune(alp, ch) {
				i := strings.IndexRune(alp, ch)
				fmt.Printf("%c", sa[i])
			} else if ch == '~' {
				ch, _, err := util.ReadRuneTTY(r, 0)
				if err != nil {
					quit = true
				}
				if ch == 3 {
					quit = true
				}
				v := "aeiou"
				if strings.ContainsRune(v, ch) {
					i := strings.IndexRune(v, ch)
					fmt.Printf("%c", svb[i])
				} else {
					fmt.Print("~")
					r.UnreadRune()
				}
			} else {
				fmt.Printf("%c", ch)
			}
		}
	}
}
