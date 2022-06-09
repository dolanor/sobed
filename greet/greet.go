package greet

import (
	"embed"
)
import "C"

//go:generate gcc -shared -fPIC -o libgreet.so ../libgreet/greet.c

//go:embed libgreet.so
var LibFS embed.FS
