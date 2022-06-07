package main

//  // #cgo LDFLAGS: -L${SRCDIR}/clib -lgreet
//  // #cgo CFLAGS: -I./clib
//  // #include "greet.h"
//  import "C"
//
//  func Greet(name string) string {
//  	cstr := C.CString(name)
//  	cstr = C.greet(cstr)
//  	s := C.GoString(cstr)
//  	return s
//  }
