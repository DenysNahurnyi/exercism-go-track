package scrabble

import (
	"strings"
)

const (
	ONE_POINT_SYMBOLS   = "aeioulnrstAEIOULNRST"
	TWO_POINT_SYMBOLS   = "dgDG"
	THREE_POINT_SYMBOLS = "bcmpBCMP"
	FOUR_POINT_SYMBOLS  = "fhvwyFHVWY"
	FIVE_POINT_SYMBOLS  = "kK"
	EIGHT_POINT_SYMBOLS = "jxJX"
	TEN_POINT_SYMBOLS   = "qzQZ"
)

func Score(str string) (points int) {
	for _, leter := range str {
		if strings.Contains(ONE_POINT_SYMBOLS, string(leter)) {
			points += 1
		}
		if strings.Contains(TWO_POINT_SYMBOLS, string(leter)) {
			points += 2
		}
		if strings.Contains(THREE_POINT_SYMBOLS, string(leter)) {
			points += 3
		}
		if strings.Contains(FOUR_POINT_SYMBOLS, string(leter)) {
			points += 4
		}
		if strings.Contains(FIVE_POINT_SYMBOLS, string(leter)) {
			points += 5
		}
		if strings.Contains(EIGHT_POINT_SYMBOLS, string(leter)) {
			points += 8
		}
		if strings.Contains(TEN_POINT_SYMBOLS, string(leter)) {
			points += 10
		}
	}
	return points
}
