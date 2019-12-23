package scale

var chromatic = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var maj = [8]int{0, 2, 2, 1, 2, 2, 2, 1}
var min = [8]int{0, 2, 1, 2, 2, 1, 2, 2}

func index(note string) int {
	for i, v := range chromatic {
		if v == note {
			return i
		}
	}
	return -1
}

// Get prints the scale for a given note
func Get(s string, note string) [7]string {
	var major = [7]string{}
	var scale = [8]int{}
	var i = index(note)

	if s == "major" {
		scale = maj
	} else if s == "minor" {
		scale = min
	}

	for j := 0; j < 7; j++ {
		i = (i + scale[j]) % 12
		major[j] = chromatic[i]
	}

	return major
}
