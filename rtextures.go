// Code generated by go generate; DO NOT EDIT.

package raylib

import (
	"unsafe"

	"github.com/judah-caruso/raylib-purego/internal"
)

// Load image from file into CPU memory (RAM)
func LoadImage(fileName string) (res Image) {
	a0 := tocstring(fileName)
	internal.LoadImage.Call(unsafe.Pointer(&res), unsafe.Pointer(&a0))
	return res
}
