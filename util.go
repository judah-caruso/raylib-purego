package raylib

import (
	"unsafe"
)

// ToString converts a C string into a Go string.
// The returned Go string points to the original C string.
func ToString(cstring *byte) string {
	if cstring == nil || *cstring == 0 {
		return ""
	}

	n := 0
	for ptr := unsafe.Pointer(cstring); *(*byte)(ptr) != 0; n += 1 {
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}

	return string(unsafe.Slice(cstring, n))
}

// ToStringN creates a new string from an array of bytes (up to the null terminator if found).
func ToStringN(cstring []byte) string {
	if cstring == nil || len(cstring) == 0 || cstring[0] == 0 {
		return ""
	}

	for i, b := range cstring {
		if b == 0 {
			return string(cstring[:i-1])
		}
	}

	return string(cstring)
}

func tocstring(gostring string) *byte {
	bytes := make([]byte, len(gostring)+1)
	copy(bytes, gostring)
	return &bytes[0]
}
