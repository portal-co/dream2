package main

import (
	"fmt"
	"io"
	"os"
	"syscall"

	shell "github.com/ipfs/go-ipfs-api"
	"portal.io/dream/util"
)

func main() {
	sh := shell.NewLocalShell()
	if os.Args[1] == "create" {
		pay := os.Args[2]
		r, w := io.Pipe()
		w.Write([]byte(`prev="$1";shift`))
		io.Copy(w, os.Stdin)
		bat, err := sh.Add(r)
		if err == nil {
			x, err := util.AddDir(sh, map[string]string{"pay": pay, "bat.sh": bat})
			if err == nil {
				fmt.Print(x)
			}
		}
	}
	if os.Args[1] == "run" {
		x := os.Args[2]
		l, err := sh.List(x)
		if err != nil {
			return
		}
		p, err := os.Getwd()
		if err != nil {
			return
		}
		for _, m := range l {
			if m.Name == "pay" {
				os.Chdir("/ipfs/" + m.Hash)
			}
		}
		for syscall.Exec("/bin/sh", append([]string{"sh", "/ipfs/" + x + "/bat.sh", p}, os.Args[3:]...), os.Environ()) != nil {

		}
	}
}
