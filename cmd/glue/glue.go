package main

import (
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"portal.io/dream/util"
)

func main() {
	sh := shell.NewLocalShell()
	m := map[string]string{}
	for _, a := range os.Args[1:] {
		s := strings.SplitN(a, ":", 2)
		m[s[0]] = s[1]
	}
	x, err := util.AddDir(sh, m)
	if err == nil {
		fmt.Print(x)
	}
}
