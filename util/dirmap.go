package util

import (
	"sync"

	shell "github.com/ipfs/go-ipfs-api"
)

type DirMap[V any] struct {
	Src   string
	Mutex sync.Mutex
}

func GetD[V any](m DirMap[V], k string) IpfsLazy[V] {
	return IpfsLazy[V]{Src: m.Src + "/" + k}
}

func PutD[V any](sh *shell.Shell, m *DirMap[V], k string, v IpfsLazy[V]) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	x, err := sh.PatchLink(m.Src, k, v.Src, true)
	if err != nil {
		return err
	}
	m.Src = x
	return nil
}

func KeysD[V any](sh *shell.Shell, x DirMap[V]) ([]string, error) {
	l, err := sh.List(x.Src)
	if err != nil {
		return nil, err
	}
	m := make([]string, len(l))
	for i, n := range l {
		m[i] = n.Name
	}
	return m, nil
}
