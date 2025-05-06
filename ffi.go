//go:generate go run ./internal/generate
package raylib

import "github.com/judah-caruso/ffi-embeded"

var (
	void = ffi.TypeVoid
	ptr  = ffi.TypePointer
	cstr = ffi.TypePointer
	ustr = ffi.TypePointer

	intPtr   = ffi.TypePointer
	uintPtr  = ffi.TypePointer
	floatPtr = ffi.TypePointer

	char  = ffi.TypeSint8
	short = ffi.TypeSint16
	cint  = ffi.TypeSint32
	long  = ffi.TypeSint64

	uchar  = ffi.TypeUint8
	ushort = ffi.TypeUint16
	ucint  = ffi.TypeUint32
	ulong  = ffi.TypeUint64

	float  = ffi.TypeFloat
	double = ffi.TypeDouble

	cbool = ffi.TypeUint8

	tRectanglePtrPtr = ptr
)
