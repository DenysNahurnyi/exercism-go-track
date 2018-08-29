// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	PROVERB    = "For want of a %s the %s was lost."
	CONCLUSION = "And all for the want of a %s."
)

// Proverb -> create proverb from template and input following some interesting rules
func Proverb(rhyme []string) (proverb []string) {
	for i := 1; i < len(rhyme); i++ {
		proverb = append(proverb, fmt.Sprintf(PROVERB, rhyme[i-1], rhyme[i]))
	}
	if len(rhyme) > 0 {
		proverb = append(proverb, fmt.Sprintf(CONCLUSION, rhyme[0]))
	}
	return proverb
}
