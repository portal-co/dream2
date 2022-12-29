package embeds

import (
	_ "embed"
)

//go:embed interp.sh
var Interp string
