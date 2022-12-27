package rat

import (
	"encoding/gob"
	"errors"
	"io"

	"github.com/awgh/bencrypt/bc"
	"github.com/awgh/ratnet/api"
	shell "github.com/ipfs/go-ipfs-api"
	"portal.io/dream/util"
)

type RatIpfs struct {
	Ipfs   *shell.Shell
	Gets   Gets
	PubKey util.IpfsLazy[bc.PubKey]
}

type Gets interface {
	Send(x string) error
	Recv() (string, error)
}

func (r *RatIpfs) RPC(host string, method api.Action, args ...interface{}) (interface{}, error) {
	switch method {
	case api.Pickup:
		x, err := r.Gets.Recv()
		if err != nil {
			return nil, err
		}
		c, err := r.Ipfs.Cat(x)
		if err != nil {
			return nil, err
		}
		var b api.Bundle
		gob.NewDecoder(c).Decode(&b)
		return b, nil
	case api.Dropoff:
		bundle := args[0].(api.Bundle)
		rr, w := io.Pipe()
		gob.NewEncoder(w).Encode(bundle)
		h, err := r.Ipfs.Add(rr)
		if err != nil {
			return nil, err
		}
		return nil, r.Gets.Send(h)
	case api.ID:
		return util.GetLazy(r.PubKey, r.Ipfs), nil
	}
	return nil, errors.New("Not Implemented")
}
