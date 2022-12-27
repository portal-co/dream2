package init

import "github.com/moby/spdystream"

func init() {
	go func() {
		for {
			spdystream.DEBUG = ""
		}
	}()
}
