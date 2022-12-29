package action

import (
	"bytes"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"portal.io/dream/embeds"
	"portal.io/dream/util"
)

type Action struct {
	Script string
	Deps   util.DirMap[Action]
}

func (a Action) Stringify(sh *shell.Shell) (string, error) {
	prelude := ""
	kk, err := util.KeysD(sh, a.Deps)
	if err != nil {
		return "", err
	}
	mm := make(map[string]*Action)
	for _, k := range kk {
		l := util.GetLazy(util.GetD(a.Deps, k), sh)
		mm[k] = l
		m, err := l.Stringify(sh)
		if err != nil {
			return "", err
		}
		prelude = fmt.Sprintf("%s %s", m, prelude)
	}
	content, err := a.Render(sh, mm)
	if err != nil {
		return "", err
	}
	return (prelude + " " + content), nil
}

func (a Action) Render(sh *shell.Shell, deps map[string]*Action) (string, error) {
	var b bytes.Buffer
	b.Write([]byte(a.Script))
	b.Write([]byte(embeds.Interp))
	return sh.Add(&b)
}
