package hamming

import "errors"

// Calculate Hamming distance between 2 genoms
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Length of genoms are not equal")
	}
	diff := 0
	for i, vA := range a {
		if vA != rune(b[i]) {
			diff++
		}
	}
	return diff, nil
}
