package scale

import (
	"fmt"
	"sort"

	"github.com/barissimsek/npp/pkg/music"
)

func index(note string) int {
	var chromatic = music.Chromatic()

	for i, v := range chromatic {
		if v == note {
			return i
		}
	}
	return -1
}

// Get returns the scale for a given note
func Get(s string, root string) []string {
	var i = index(root)
	var formula = make([]int, 8)
	var scale = []string{}

	var chromatic = music.Chromatic()

	// Choose interval formula
	formula = music.Scale(s)

	// Generate scale from intervals
	for j := 0; j < 7; j++ {
		i = (i + formula[j]) % 12
		scale = append(scale, chromatic[i])
	}

	return scale
}

// List returns major and minor scales for all notes
func List() map[string][]string {
	var m = make(map[string][]string)
	var r []string

	for _, v := range music.Chromatic() {
		r = append(r, v+"maj")
		r = append(r, v+"min")
		m[v+"maj"] = Get("maj", v)
		m[v+"min"] = Get("minn", v)
	}

	sort.Strings(r)
	for _, k := range r {
		fmt.Printf("%10s: %s\n", k, m[k])
	}

	return m
}
