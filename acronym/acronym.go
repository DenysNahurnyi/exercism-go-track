// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
)

func splitter(str string) []string {
	re := regexp.MustCompile(` |,|-`)
	strs := re.Split(str, -2)
	for i, v := range strs {
		if v == " " {
			strs = append(strs[:i], strs[i+1:]...)
		}
	}
	return strs
}

func Abbreviate(s string) (abbr string) {
	words := splitter(s)
	for _, v := range words {
		if len(v) > 0 {
			if v[0] >= 97 && v[0] <= 122 {
				abbr += string(v[0] - 32)
			} else {
				abbr += string(v[0])
			}
		}
	}
	return abbr
}
