package raylib

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"unsafe"

	"github.com/judah-caruso/ffi-embeded"
)

var (
	void = ffi.TypeVoid
	ptr  = ffi.TypePointer
	cstr = ffi.TypePointer

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
)

func init() {
	runtime.LockOSThread()
}

type cproc struct {
	addr uintptr
	cif  ffi.Cif
}

func (c *cproc) call(res unsafe.Pointer, args ...unsafe.Pointer) {
	ffi.Call(&c.cif, c.addr, res, args...)
}

func (c *cproc) void(args ...unsafe.Pointer)          { c.call(nil, args...) }
func (c *cproc) bool(args ...unsafe.Pointer) bool     { return c.uint8(args...) != 0 }
func (c *cproc) uint8(args ...unsafe.Pointer) uint8   { return uint8(c.uint64(args...)) }
func (c *cproc) uint16(args ...unsafe.Pointer) uint16 { return uint16(c.uint64(args...)) }
func (c *cproc) uint32(args ...unsafe.Pointer) uint32 { return uint32(c.uint64(args...)) }
func (c *cproc) uint64(args ...unsafe.Pointer) (res uint64) {
	c.call(unsafe.Pointer(&res), args...)
	return
}

func (c *cproc) int8(args ...unsafe.Pointer) int8   { return int8(c.uint64(args...)) }
func (c *cproc) int16(args ...unsafe.Pointer) int16 { return int16(c.uint64(args...)) }
func (c *cproc) int32(args ...unsafe.Pointer) int32 { return int32(c.uint64(args...)) }
func (c *cproc) int64(args ...unsafe.Pointer) int64 { return int64(c.uint64(args...)) }

var raylib ffi.Lib

func prep(retType *ffi.Type, name string, argTypes ...*ffi.Type) cproc {
	// Lazily load the library
	if raylib.Addr == 0 {
		name := fmt.Sprintf("%d-%s", time.Now().UnixNano(), libName)
		path := filepath.Join(os.TempDir(), name)

		var err error
		if err = os.WriteFile(path, libData, 0644); err != nil {
			panic(err)
		}

		defer os.Remove(path)

		raylib, err = ffi.Load(path)
		if err != nil {
			panic(err)
		}
	}

	var cif ffi.Cif
	if status := ffi.PrepCif(&cif, ffi.DefaultAbi, uint32(len(argTypes)), retType, argTypes...); status != ffi.OK {
		panic(fmt.Sprintf("%s: %v", name, status))
	}

	procAddress, err := raylib.Get(name)
	if err != nil {
		panic(err)
	}

	return cproc{
		addr: procAddress,
		cif:  cif,
	}
}
