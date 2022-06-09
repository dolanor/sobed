package main

// #cgo LDFLAGS: -ldl
// #include "dlfcn.h"
// char*
// greetWrap(void *f, char *name)
// {
//   char* (*greetfn)(char *);
//
//   greetfn = f;
//   return greetfn(name);
// }
import "C"
import (
	"embed"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"unsafe"

	"github.com/dolanor/sobed/greet"
)

type CGreeter struct {
	greetFn unsafe.Pointer
}

func (g *CGreeter) Greet(name string) string {
	cstr := C.CString(name)
	cstr, err := C.greetWrap(g.greetFn, cstr)
	if err != nil {
		log.Panic("error:", err)
	}
	s := C.GoString(cstr)
	return s
}

// dislodgeAndDLOpen will dislodge libName from embedFS, save it as a temp file on the system.
// From there, it will dlopen it, and get the funcName from it.
func dislodgeAndDLOpen(embedFS embed.FS, libName, funcName string) (unsafe.Pointer, error) {
	log.Println("creating temp dir to host .so")
	dirName, err := ioutil.TempDir("", "*")
	if err != nil {
		return nil, err
	}

	log.Println("opening a .so from embed")
	f, err := greet.LibFS.Open(libName)
	if err != nil {
		return nil, err
	}

	log.Println("read .so content into memory buffer")
	b := make([]byte, 1_000_000)
	_, err = f.Read(b)
	if err != nil {
		return nil, err
	}

	fPath := filepath.Join(dirName, libName)

	log.Println("writing embedded .so in:", fPath)
	err = ioutil.WriteFile(fPath, b, 0o600)
	if err != nil {
		return nil, err
	}

	log.Println("dlopen the lib apps from temp file")
	lName := C.CString(fPath)
	handle := C.dlopen(lName, C.RTLD_LAZY)
	if handle == nil {
		return nil, errors.New("dlopen: fail")
	}

	log.Println("dlsym the func from the lib")
	symbol := C.CString(funcName)
	greetFn := C.dlsym(handle, symbol)
	if greetFn == nil {
		return nil, errors.New("dlsym: fail")
	}
	return greetFn, err
}
