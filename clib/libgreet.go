package greet

import "C"
import _ "embed"

//go:embed libgreet.so
var LibGreet []byte
