package scale

var chromatic = []string{"A", "A#/Bb", "B", "C", "C#/Db", "D", "D#/Eb", "E", "F", "F#/Gb", "G", "G#/Ab"}
var scalemap = map[string][]int{
	"maj":  []int{0, 2, 2, 1, 2, 2, 2, 1}, // Major intervals starting from the root note
	"minn": []int{0, 2, 1, 2, 2, 1, 2, 2}, // Natural minor intervals starting from the root note
	"minh": []int{0, 2, 1, 2, 2, 1, 2, 2}, // Harmonic minor intervals starting from the root note
	"minm": []int{0, 2, 1, 2, 2, 1, 2, 2}, // Melodic minor intervals starting from the root note
}

func index(note string) int {
	for i, v := range chromatic {
		if v == note {
			return i
		}
	}
	return -1
}

// Get returns the scale for a given note
func Get(s string, note string) []string {
	var i = index(note)
	var formula = make([]int, 8)
	var scale = []string{}

	// Choose interval formula
	if _, ok := scalemap[s]; ok {
		formula = scalemap[s]
	} else {
		return nil
	}

	// Generate scale from intervals
	for j := 0; j < 7; j++ {
		i = (i + formula[j]) % 12
		scale = append(scale, chromatic[i])
	}

	return scale
}
