package scrabble

import (
	"strings"
)

const (
	onePointSymbols   = "aeioulnrstAEIOULNRST"
	twoPointSymbols   = "dgDG"
	threePointSymbols = "bcmpBCMP"
	fourPointSymbols  = "fhvwyFHVWY"
	fivePointSymbols  = "kK"
	eightPointSymbols = "jxJX"
	tenPointSymbols   = "qzQZ"
)

// Return amount of scores for a string with symbols
func Score(str string) (points int) {
	for _, leterRaw := range str {
		leter := string(leterRaw)
		if strings.Contains(onePointSymbols, leter) {
			points += 1
			continue
		}
		if strings.Contains(twoPointSymbols, leter) {
			points += 2
			continue
		}
		if strings.Contains(threePointSymbols, leter) {
			points += 3
			continue
		}
		if strings.Contains(fourPointSymbols, leter) {
			points += 4
			continue
		}
		if strings.Contains(fivePointSymbols, leter) {
			points += 5
			continue
		}
		if strings.Contains(eightPointSymbols, leter) {
			points += 8
			continue
		}
		if strings.Contains(tenPointSymbols, leter) {
			points += 10
			continue
		}
	}
	return points
}
