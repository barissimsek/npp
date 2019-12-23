package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/barissimsek/npp/scale"
)

func main() {
	var note = flag.String("note", "C", "Root note for the scale")
	var s = flag.String("scale", "major", "scale: major, minor")
	flag.Parse()

	major := scale.Get(*s, *note)

	fmt.Println(major)

	os.Exit(0)
}
