package luhn

import (
	"strconv"
	"strings"
)

func Valid(idStr string) bool {
	idStr = strings.Replace(idStr, " ", "", -1)
	if len(idStr) < 2 {
		return false
	}
	sum := 0
	for i := len(idStr) - 1; i > -1; i-- {
		num, err := strconv.Atoi(string(idStr[i]))
		if err != nil {
			return false
		}
		if (len(idStr)-i+1)%2 == 1 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	return sum%10 == 0
}
