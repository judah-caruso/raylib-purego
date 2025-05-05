package raylib

import (
	"unsafe"
)

func tocstring(gostring string) *byte {
	bytes := make([]byte, len(gostring)+1)
	copy(bytes, gostring)
	return &bytes[0]
}

func tostring(cstring *byte) string {
	if cstring == nil || *cstring == 0 {
		return ""
	}

	n := 0
	for ptr := unsafe.Pointer(cstring); *(*byte)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}

	return string(unsafe.Slice(cstring, n))
}
