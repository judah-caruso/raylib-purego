package internal

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/judah-caruso/ffi-embeded"
)

func init() {
	runtime.LockOSThread()
}

type cproc struct {
	addr uintptr
	cif  ffi.Cif
}

func (c *cproc) Call(res unsafe.Pointer, args ...unsafe.Pointer) {
	ffi.Call(&c.cif, c.addr, res, args...)
}

func (c *cproc) Void(args ...unsafe.Pointer)          { c.Call(nil, args...) }
func (c *cproc) Bool(args ...unsafe.Pointer) bool     { return c.Uint8(args...) != 0 }
func (c *cproc) Uint8(args ...unsafe.Pointer) uint8   { return uint8(c.Uint64(args...)) }
func (c *cproc) Uint16(args ...unsafe.Pointer) uint16 { return uint16(c.Uint64(args...)) }
func (c *cproc) Uint32(args ...unsafe.Pointer) uint32 { return uint32(c.Uint64(args...)) }
func (c *cproc) Uint64(args ...unsafe.Pointer) (res uint64) {
	c.Call(unsafe.Pointer(&res), args...)
	return
}

func (c *cproc) Int8(args ...unsafe.Pointer) int8   { return int8(c.Uint64(args...)) }
func (c *cproc) Int16(args ...unsafe.Pointer) int16 { return int16(c.Uint64(args...)) }
func (c *cproc) Int32(args ...unsafe.Pointer) int32 { return int32(c.Uint64(args...)) }
func (c *cproc) Int64(args ...unsafe.Pointer) int64 { return int64(c.Uint64(args...)) }

var raylib ffi.Lib

func Bind(retType *ffi.Type, name string, argTypes ...*ffi.Type) cproc {
	// Lazily load the library
	if raylib.Addr == 0 {
		file, err := os.CreateTemp("", "*-"+LibraryName)
		if err != nil {
			panic(err)
		}

		path := file.Name()
		libBytes := unsafe.Slice(unsafe.StringData(LibraryData), len(LibraryData))
		if _, err = file.Write(libBytes); err != nil {
			panic(err)
		}

		if err = file.Close(); err != nil {
			panic(err)
		}

		raylib, err = ffi.Load(path)
		if err != nil {
			panic(err)
		}

		os.Remove(path)
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
