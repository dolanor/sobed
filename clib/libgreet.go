package greet

import (
	"embed"
)
import "C"

//go:embed libgreet.so
var LibGreet embed.FS
