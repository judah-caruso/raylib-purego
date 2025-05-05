//go:build unix && !darwin

package raylib

import _ "embed"

const libName = "libraylib.so.5.5.0"

//go:embed libraylib.so.5.5.0
var libData []byte
