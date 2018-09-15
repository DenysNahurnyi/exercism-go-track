package isogram

func isLower(s rune) bool {
	return byte(s) >= 97 && byte(s) <= 122
}

func isCapital(s rune) bool {
	return byte(s) >= 65 && byte(s) <= 90
}

func isLetter(s rune) bool {
	return isCapital(s) || isLower(s)
}

// IsISogram checks whethere `str` string is isogram (set of letter with all letters unique)
func IsIsogram(str string) bool {
	lettersSet := make(map[rune]bool)
	for _, v := range str {
		if !isLetter(v) {
			continue
		}
		if isLower(v) {
			v = rune(int(v) - 32)
		}
		if lettersSet[v] {
			return false
		}
		lettersSet[v] = true
	}
	return true
}
