//go:build darwin

package raylib

import _ "embed"

const libName = "libraylib.5.5.0.dylib"

//go:embed libraylib.5.5.0.dylib
var libData []byte
