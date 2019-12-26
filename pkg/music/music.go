package music

var chromatic = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var scalemap = map[string][]int{
	"maj":  []int{0, 2, 2, 1, 2, 2, 2, 1}, // Major intervals starting from the root note
	"minn": []int{0, 2, 1, 2, 2, 1, 2, 2}, // Natural minor intervals starting from the root note
	"minh": []int{0, 2, 1, 2, 2, 1, 2, 2}, // Harmonic minor intervals starting from the root note
	"minm": []int{0, 2, 1, 2, 2, 1, 2, 2}, // Melodic minor intervals starting from the root note
}

// Chromatic returns chromatic scale
func Chromatic() []string {
	return chromatic
}

// Scale returns the scale for a given mode
func Scale(mode string) []int {
	if _, ok := scalemap[mode]; ok {
		return scalemap[mode]
	}

	return nil
}
