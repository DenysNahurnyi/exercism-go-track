// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

func isLetter(s rune) bool {
	return isCapital(s) || isLower(s)
}
func isNumber(s rune) bool {
	return byte(s) >= 48 && byte(s) <= 57
}
func isCapital(s rune) bool {
	return byte(s) >= 65 && byte(s) <= 90
}
func isLower(s rune) bool {
	return byte(s) >= 97 && byte(s) <= 122
}
func isSymbol(s rune) bool {
	return (byte(s) > 31 && byte(s) < 65) || (byte(s) > 90 && byte(s) < 97) || byte(s) > 122
}
func isSpace(s rune) bool {
	return byte(s) == 32
}

func isShout(str string) bool {
	var letters, capitals int
	for _, v := range str {
		if isLetter(v) {
			letters++
			if isCapital(v) {
				capitals++
			}
		}
	}
	return letters == capitals && letters > 0
}

func isQuestion(str string) bool {
	var resp bool
	for _, v := range str {
		if strings.Contains(string(v), "?") {
			resp = true
		}
		if isLetter(v) {
			resp = false
		}
	}
	return resp
}

func isWithoutWords(str string) bool {
	for _, v := range str {
		if isLetter(v) || isNumber(v) {
			return false
		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {

	if isQuestion(remark) {
		if isShout(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if isWithoutWords(remark) {
		return "Fine. Be that way!"
	}
	if isShout(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
