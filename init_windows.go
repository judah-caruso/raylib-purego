//go:build windows

package raylib

import _ "embed"

const libName = "raylib.dll"

//go:embed raylib.dll
var libData []byte
