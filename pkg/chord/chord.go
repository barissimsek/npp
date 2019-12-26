package chord

import (
	"fmt"
	"sort"

	"github.com/barissimsek/npp/pkg/music"
	"github.com/barissimsek/npp/pkg/scale"
)

// Get returns triads for the given chord
func Get(t string, root string) [3]string {
	var s []string
	var c [3]string

	s = scale.Get(t, root)

	c[0] = s[0]
	c[1] = s[2]
	c[2] = s[4]

	return c
}

// List returns all chords and triads
func List() map[string][3]string {
	var m = make(map[string][3]string)
	var c []string

	for _, root := range music.Chromatic() {
		c = append(c, root+"maj")
		c = append(c, root+"min")
		m[root+"maj"] = Get("maj", root)
		m[root+"min"] = Get("minn", root)
	}

	sort.Strings(c)
	for _, k := range c {
		fmt.Printf("%10s: %s\n", k, m[k])
	}

	return m
}
