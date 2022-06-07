package main

// #cgo LDFLAGS: -ldl
// #include "dlfcn.h"
// char*
// my_greet(void *f, char *name)
// {
//   char* (*greetfn)(char *);
//
//   greetfn = f;
//   return greetfn(name);
// }
import "C"
import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	greet "github.com/dolanor/sobed/clib"
)

func load() {
	libName := "libgreet.so"
	dirName, err := ioutil.TempDir("", "*")
	if err != nil {
		log.Fatal(err)
	}

	fPath := filepath.Join(dirName, libName)
	f, err := greet.LibGreet.Open(libName)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 1_000_000)
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("data:", string(b))

	fi, err := f.Stat()
	// NOT POSSIBLE TO GET FD FROM EMBED
	if err != nil {
		log.Fatal(err)
	}
	_ = fi

	//fi.Sys
	//st := fi.Sys()

	//log.Printf("%T - %+v\n", fi, fi)
	time.Sleep(30 * time.Second)
	return

	//err = ioutil.WriteFile(fPath, greet.LibGreet, 0o600)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("written in:", fPath)

	lName := C.CString(fPath)
	handle := C.dlopen(lName, C.RTLD_LAZY)
	if handle == nil {
		log.Println("dammit: nil")
		return
	}

	symbol := C.CString("greet")
	greetFn := C.dlsym(handle, symbol)
	if greetFn == nil {
		log.Println("dammit: sym nil")
		return
	}

	name := "tato"
	cstr := C.CString(name)
	cstr, err = C.my_greet(greetFn, cstr)
	if err != nil {
		log.Panic("MDR:", err)
	}
	s := C.GoString(cstr)
	fmt.Println("CACA:", s)
}
