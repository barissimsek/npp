package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/barissimsek/npp/pkg/scale"
)

// Usage: cmd -scale major -note C
func main() {
	var note = flag.String("note", "C", "Root note for the scale")
	var s = flag.String("scale", "major", "scale: major, minor")
	flag.Parse()

	scale := scale.Get(*s, *note)

	fmt.Println(scale)

	os.Exit(0)
}
