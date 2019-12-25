package chord

import (
	"github.com/barissimsek/npp/pkg/scale"
)

// Get returns major triad chord for the given note
func Get(t string, note string) [3]string {
	var s []string
	var c [3]string

	s = scale.Get(t, note)

	c[0] = s[0]
	c[1] = s[2]
	c[2] = s[4]

	return c
}
